package controllertests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Victoria-engine/api-v2/app/utils/testutils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestMain : TestMain
func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	testutils.Database()

	os.Exit(m.Run())
}

func TestRegister(t *testing.T) {
	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = testutils.SeedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	requests := []struct {
		inputJSON    string
		statusCode   int
		email        string
		password     string
		errorMessage string
	}{
		// Expected to PASS
		{
			inputJSON:    `{"email": "asd@email.com", "password":"asd", "first_name": "Tiago", "last_name":"Ferreira"}`,
			statusCode:   200,
			errorMessage: "",
		},
		// Expected to FAIL - Missing data
		{
			inputJSON:    `{"email": "asd@email.com", "password":"asd"}`,
			statusCode:   422,
			errorMessage: "Required First Name",
		},
		// Expected to FAIL - Missing data
		{
			inputJSON:    `{"email": "asd@email.com", "password":"asd", "first_name":"Tiago"}`,
			statusCode:   422,
			errorMessage: "Required Last Name",
		},
		// Expected to FAIL - Duplicated record
		{
			inputJSON:    `{"email": "asd@email.com", "password":"asd", "first_name": "Tester", "last_name":"123"}`,
			statusCode:   422,
			errorMessage: "User email already exists",
		},
	}

	for _, reqData := range requests {
		req, err := http.NewRequest("POST", "api/auth/register", bytes.NewBufferString(reqData.inputJSON))
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(testutils.Server.Register)
		handler.ServeHTTP(rr, req)

		// Test the status code
		assert.Equal(t, rr.Code, reqData.statusCode)
		if reqData.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		// Test the error message
		if reqData.statusCode == 422 && reqData.errorMessage != "" {
			responseMap := make(map[string]interface{})

			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}

			assert.Equal(t, responseMap["error"], reqData.errorMessage)
		}
	}
}

func TestLogin(t *testing.T) {
	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = testutils.SeedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	requests := []struct {
		inputJSON    string
		statusCode   int
		email        string
		password     string
		errorMessage string
	}{
		// Expected to PASS
		{
			inputJSON:    `{"email": "logged@email.com", "password":"password"}`,
			statusCode:   200,
			errorMessage: "",
		},
		// Expected to FAIL - Wrong data
		{
			inputJSON:    `{"email": "logged@email.com", "password":"asd"}`,
			statusCode:   422,
			errorMessage: "Incorrect Password",
		},
		// Expected to FAIL - Missing data
		{
			inputJSON:    `{"email": "logged@email.com"}`,
			statusCode:   422,
			errorMessage: "Required Password",
		},
		// Expected to FAIL - Missing data
		{
			inputJSON:    `{"password": "password"}`,
			statusCode:   422,
			errorMessage: "Required Email",
		},
		// Expected to FAIL - Record not found
		{
			inputJSON:    `{"email": "asd@email.com",  "password":"password"}`,
			statusCode:   422,
			errorMessage: "Record does not exist",
		},
	}

	for _, reqData := range requests {
		req, err := http.NewRequest("POST", "api/auth/register", bytes.NewBufferString(reqData.inputJSON))
		if err != nil {
			t.Error(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(testutils.Server.Login)
		handler.ServeHTTP(rr, req)

		// Test the status code
		assert.Equal(t, rr.Code, reqData.statusCode)
		if reqData.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		// Test the error message
		if reqData.statusCode == 422 && reqData.errorMessage != "" {
			responseMap := make(map[string]interface{})

			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to json: %v", err)
			}

			assert.Equal(t, responseMap["error"], reqData.errorMessage)
		}
	}
}
