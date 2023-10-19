package service

import (
	"testing"

	"github.com/sinisaos/chi-ent/ent/enttest"
	"github.com/sinisaos/chi-ent/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestAnswerService(t *testing.T) {
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
	// Insert question
	q, err := questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "TestQuestion1",
		Author:  u.ID,
		Content: "Content of question one",
	})
	assert.NoError(t, err)

	// Insert answers
	a, err := answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer one",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.Equal(t, a.ID, 1)
	assert.NoError(t, err)

	_, err = answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer two",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.NoError(t, err)

	// Check if page and itemsPerPage works
	itemsPerPageOne, _ := answerService.GetAllAnswers(1, 1)
	assert.Len(t, itemsPerPageOne, 1)
	itemsPerPageTwo, _ := answerService.GetAllAnswers(1, 2)
	assert.Len(t, itemsPerPageTwo, 2)
	secondPageResult, _ := answerService.GetAllAnswers(2, 1)
	assert.Len(t, secondPageResult, 1)

	// Check empty data result
	emptyDataResult, _ := answerService.GetAllAnswers(2, 4)
	assert.Len(t, emptyDataResult, 0)

	// Single Answer
	resultSingleAnswer, _ := answerService.GetAnswer(1)
	assert.Contains(t, resultSingleAnswer.Content, "Content of Answer one")

	// Return error if answer does not exist
	_, err = answerService.GetAnswer(10)
	assert.Error(t, err)

	// Update answer if exist
	_, err = answerService.UpdateAnswer(2, &model.UpdateAnswerInput{
		Content: "Updated content od Answer two",
	})
	assert.NoError(t, err)

	// Check error if answer does not exists
	_, err = answerService.UpdateAnswer(3, &model.UpdateAnswerInput{
		Content: "Updated content od Answer three",
	})
	assert.Error(t, err)

	// Delete answer
	err = answerService.DeleteAnswer(2)
	assert.NoError(t, err)

	// Check error if answer does not exists
	err = answerService.DeleteAnswer(10)
	assert.Error(t, err)

	// Check if answer is deleted
	deletedAnswerResult, _ := answerService.GetAllAnswers(1, 1)
	assert.Len(t, deletedAnswerResult, 1)

	// Checking answer author and question
	answerAuthor, _ := answerService.GetAnswerAuthor(a.ID)
	assert.Equal(t, answerAuthor.Edges.Author.ID, 1)
	assert.Equal(t, answerAuthor.Edges.Author.Username, "TestUser1")

	answerQuestion, _ := answerService.GetAnswerQuestion(a.ID)
	assert.Equal(t, answerQuestion.Edges.Question.ID, 1)
	assert.Equal(t, answerQuestion.Edges.Question.Title, "TestQuestion1")
}
