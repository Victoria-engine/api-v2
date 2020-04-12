package modeltests

import (
	"log"
	"testing"

	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/Victoria-engine/api-v2/app/utils/testutils"
	"github.com/stretchr/testify/assert"
)

// [User] Save
func TestSaveUser(t *testing.T) {

	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@email.com",
		Password:  "asd",
	}

	user.Prepare()

	createdUser, err := user.SaveUser(testutils.Server.DB)
	if err != nil {
		t.Errorf("Error saving the user: %v\n", err)
		return
	}

	assert.Equal(t, user.ID, createdUser.ID)
	assert.Equal(t, createdUser.BlogID, -1)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.FirstName, createdUser.FirstName)
	assert.Equal(t, user.LastName, createdUser.LastName)
}

func TestFindUserByID(t *testing.T) {
	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := testutils.SeedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	foundUser, err := testutils.UserInstance.FindUserByID(testutils.Server.DB, user.ID)
	if err != nil {
		t.Errorf("Got error trying to find one user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
}

// [User] Delete
func TestUserDeleteByID(t *testing.T) {
	err := testutils.RefreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := testutils.SeedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := testutils.UserInstance.DeleteByID(testutils.Server.DB, user.ID)
	if err != nil {
		t.Errorf("Got an error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
