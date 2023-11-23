package service

import (
	"testing"

	"github.com/gosimple/slug"
	"github.com/sinisaos/chi-ent/ent/enttest"
	"github.com/sinisaos/chi-ent/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestQuestionService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	userService := NewUserService(client)
	questionService := NewQuestionService(client)
	// tagService := NewTagService(client)
	answerService := NewAnswerService(client)
	// Insert user
	u, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	// Tags slice
	var tags []string
	tags = append(tags, "TestTag1", "TestTag2")

	// Insert question
	q, err := questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "Test Question 1",
		Slug:    slug.Make("Test Question 1"),
		Content: "Content od question one",
		Author:  u.ID,
		Tags:    tags,
	})
	assert.Equal(t, q.ID, 1)
	assert.NoError(t, err)

	_, err = questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "Test Question 2",
		Slug:    slug.Make("Test Question 2"),
		Content: "Content od question two",
		Author:  u.ID,
		Tags:    tags,
	})
	assert.NoError(t, err)

	// Check if page and itemsPerPage works
	itemsPerPageOne, _ := questionService.GetAllQuestions(1, 1)
	assert.Len(t, itemsPerPageOne, 1)
	itemsPerPageTwo, _ := questionService.GetAllQuestions(1, 2)
	assert.Len(t, itemsPerPageTwo, 2)
	secondPageResult, _ := questionService.GetAllQuestions(2, 1)
	assert.Len(t, secondPageResult, 1)

	// Check empty data result
	emptyDataResult, _ := questionService.GetAllQuestions(2, 4)
	assert.Len(t, emptyDataResult, 0)

	// Single question
	resultSingleQuestion, _ := questionService.GetQuestion(1)
	assert.Contains(t, resultSingleQuestion.Title, "Test Question 1")

	// Return error if Question does not exist
	_, err = questionService.GetQuestion(10)
	assert.Error(t, err)

	// Update tags slice
	var updatedTags []string
	updatedTags = append(updatedTags, "TestTag3", "TestTag4")

	// Update question if exist
	_, err = questionService.UpdateQuestion(2, &model.UpdateQuestionInput{
		Title:   "Test Question2 Updated",
		Content: "Updated content od question two",
		Tags:    updatedTags,
	})
	assert.NoError(t, err)

	// Check error if question does not exists
	_, err = questionService.UpdateQuestion(3, &model.UpdateQuestionInput{
		Title:   "TestQuestion3Updated",
		Content: "Updated content od question three",
		Tags:    updatedTags,
	})
	assert.Error(t, err)

	// Delete question
	err = questionService.DeleteQuestion(2)
	assert.NoError(t, err)

	// Check error if question does not exists
	err = questionService.DeleteQuestion(10)
	assert.Error(t, err)

	// Check if question is deleted
	deletedQuestionResult, _ := questionService.GetAllQuestions(1, 1)
	assert.Len(t, deletedQuestionResult, 1)

	// Insert answer for checking question answers
	_, err = answerService.CreateAnswer(&model.NewAnswerInput{
		Content:  "Content of Answer one",
		Author:   u.ID,
		Question: q.ID,
	})
	assert.NoError(t, err)

	// Checking question author, tags and answers
	questionAuthor, _ := questionService.GetQuestionAuthor(q.ID)
	assert.Equal(t, questionAuthor.Edges.Author.ID, 1)
	assert.Equal(t, questionAuthor.Edges.Author.Username, "TestUser1")

	questionAnswers, _ := questionService.GetQuestionAnswers(q.ID)
	assert.Equal(t, questionAnswers.Edges.Answers[0].ID, 1)
	assert.Equal(t, questionAnswers.Edges.Answers[0].Content, "Content of Answer one")
	assert.Len(t, questionAnswers.Edges.Answers, 1)

	questionTags, _ := questionService.GetQuestionTags(q.ID)
	assert.Equal(t, questionTags.Edges.Tags[0].ID, 1)
	assert.Equal(t, questionTags.Edges.Tags[0].Name, "TestTag1")
	assert.Len(t, questionTags.Edges.Tags, 2)
}
