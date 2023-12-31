package service

import (
	"testing"

	"github.com/sinisaos/chi-ent/ent/enttest"
	"github.com/sinisaos/chi-ent/pkg/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	userService := NewUserService(client)
	questionService := NewQuestionService(client)
	answerService := NewAnswerService(client)
	// Insert user
	u, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	_, err = userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser2",
		Email:    "testuser2@gmail.com",
		Password: "pass1234",
	})
	assert.NoError(t, err)

	// Check duplicate email error
	_, err = userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser3",
		Email:    "testuser2@gmail.com",
		Password: "pass12345",
	})
	assert.Error(t, err)

	// Check if page and itemsPerPage works
	itemsPerPageOne, _ := userService.GetAllUsers(1, 1)
	assert.Len(t, itemsPerPageOne, 1)
	itemsPerPageTwo, _ := userService.GetAllUsers(1, 2)
	assert.Len(t, itemsPerPageTwo, 2)
	secondPageResult, _ := userService.GetAllUsers(2, 1)
	assert.Len(t, secondPageResult, 1)

	// Check empty data result
	emptyDataResult, _ := userService.GetAllUsers(2, 4)
	assert.Len(t, emptyDataResult, 0)

	// Single user
	resultSingleUser, _ := userService.GetUser(1)
	assert.Contains(t, resultSingleUser.Username, "TestUser1")

	// Return error if user does not exist
	_, err = userService.GetUser(10)
	assert.Error(t, err)

	// Update user if exist
	_, err = userService.UpdateUser(2, &model.UpdateUserInput{
		UserName: "TestUser2Updated",
		Password: "pass1234updated",
	})
	assert.NoError(t, err)

	// Check error if user does not exists
	_, err = userService.UpdateUser(3, &model.UpdateUserInput{
		UserName: "TestUser2Updated",
		Password: "pass1234updated",
	})
	assert.Error(t, err)

	// Delete user
	err = userService.DeleteUser(2)
	assert.NoError(t, err)

	// Check error if user does not exists
	err = userService.DeleteUser(10)
	assert.Error(t, err)

	// Check if user is deleted
	deletedUserResult, _ := userService.GetAllUsers(1, 1)
	assert.Len(t, deletedUserResult, 1)

	// Insert question for checking user questions
	q, err := questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion1",
		Author:  u.ID,
		Content: "Content of question one",
	})
	assert.NoError(t, err)

	// Insert answer for checking user answers
	a, err := answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer one",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.Equal(t, a.ID, 1)
	assert.NoError(t, err)

	// Checking user questions and answers
	userQuestions, _ := userService.GetUserQuestions(u.ID)
	assert.Equal(t, userQuestions.Edges.Questions[0].ID, 1)
	assert.Equal(t, userQuestions.Edges.Questions[0].Title, "TestQuestion1")
	assert.Len(t, userQuestions.Edges.Questions, 1)

	userAnswers, _ := userService.GetUserAnswers(u.ID)
	assert.Equal(t, userAnswers.Edges.Answers[0].ID, 1)
	assert.Equal(t, userAnswers.Edges.Answers[0].Content, "Content of Answer one")
	assert.Len(t, userAnswers.Edges.Answers, 1)

}
