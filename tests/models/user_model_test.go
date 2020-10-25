package models_test

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luk3skyw4lker/social-go/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshUserTable()

	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()

	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)

	if err != nil {
		t.Errorf("Error on FindAllUsers: %v", err)

		return
	}

	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {
	err := refreshUserTable()

	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
		ID:       1,
		Nickname: "Lucas",
		Email:    "lucas@lucas.com",
		Password: "lucas",
	}

	savedUser, err := user.SaveUser(server.DB)

	if err != nil {
		t.Errorf("Error on save user: %v", err)

		return
	}

	assert.Equal(t, user.Nickname, savedUser.Nickname)
	assert.Equal(t, user.Email, savedUser.Email)
	assert.Equal(t, user.ID, savedUser.ID)
}

func TestGetUserById(t *testing.T) {
	err := refreshUserTable()

	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatal(err)
	}

	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)

	if err != nil {
		t.Errorf("Error on get user by id: %v", err)

		return
	}

	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T) {
	err := refreshUserTable()

	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatal(err)
	}

	userUpdate := models.User{
		ID:       1,
		Nickname: "update",
		Email:    "update@update.com",
		Password: "pass",
	}

	updatedUser, err := userUpdate.UpdateUser(server.DB, user.ID)

	if err != nil {
		t.Errorf("Error on update user: %v", err)

		return
	}

	assert.Equal(t, userUpdate.Nickname, updatedUser.Nickname)
	assert.Equal(t, userUpdate.Email, updatedUser.Email)
	assert.Equal(t, userUpdate.ID, updatedUser.ID)
}

func TestDeleteAUser(t *testing.T) {
	err := refreshUserTable()

	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed users: %v", err)
	}

	isDeleted, err := userInstance.DeleteUser(server.DB, user.ID)

	if err != nil {
		t.Errorf("Error on delete a user: %v", err)
	}

	assert.Equal(t, isDeleted, int64(1))
}
