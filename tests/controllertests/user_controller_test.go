package controllertests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Victoria-engine/api-v2/pkg/utl/testutils"
	"github.com/stretchr/testify/assert"
)

func TestGetUserInfo(t *testing.T) {
	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = testutils.SeedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	testCases := []struct {
		userID       string
		statusCode   int
		errorMessage string
	}{
		// Expected to PASS
		{
			userID:       "1",
			statusCode:   200,
			errorMessage: "",
		},
		// Expected to FAIL - Record not found
		{
			userID:       "999",
			statusCode:   404,
			errorMessage: "record not found",
		},
		// Expected to FAIL - Missing paramenter
		{
			userID:       "",
			statusCode:   400,
			errorMessage: " ",
		},
	}

	for _, reqData := range testCases {
		params := map[string]string{
			"id": reqData.userID,
		}

		rr := testutils.AuthenticatedRequest(testutils.Server.GetUserInfo, "GET", "api/users", params, nil)

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
