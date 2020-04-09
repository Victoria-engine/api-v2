package modeltests

import (
	"log"
	"testing"

	"github.com/Victoria-engine/api-v2/app/models"
	"github.com/stretchr/testify/assert"
)

// [User] Save
func TestSaveUser(t *testing.T) {

	err := refreshUsersAndPostsTable()
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

	createdUser, err := user.SaveUser(server.DB)
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
	err := refreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatal(err)
	}

	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("Got error trying to find one user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
}

// [User] Delete
func TestUserDeleteByID(t *testing.T) {
	err := refreshUsersAndPostsTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("Got an error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
