package service

import (
	"testing"

	"github.com/sinisaos/chi-ent/ent/enttest"
	"github.com/sinisaos/chi-ent/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestAuthService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	userService := NewUserService(client)
	authService := NewAuthService(client)
	// Insert user
	_, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	// Single user
	resultSingleUser, _ := userService.GetUser(1)
	assert.Contains(t, resultSingleUser.Username, "TestUser1")

	// Login user
	_, err = authService.Login(&model.LoginUserInput{
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	// Login wrong user
	_, err = authService.Login(&model.LoginUserInput{
		Email:    "wronguser@gmail.com",
		Password: "wrongpass",
	})
	assert.Error(t, err)
}
