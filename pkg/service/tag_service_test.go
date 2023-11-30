package service

import (
	"testing"

	"github.com/gosimple/slug"
	"github.com/sinisaos/chi-ent/ent/enttest"
	"github.com/sinisaos/chi-ent/pkg/model"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestTagService(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:test?mode=memory&_fk=1")
	defer client.Close()
	tagService := NewTagService(client)
	userService := NewUserService(client)
	questionService := NewQuestionService(client)

	// Insert user for checking tags questions
	u, err := userService.CreateUser(&model.NewUserInput{
		UserName: "TestUser1",
		Email:    "testuser1@gmail.com",
		Password: "pass123",
	})
	assert.NoError(t, err)

	// Tags as slice of strings
	var tags []string
	tags = append(tags, "TestTag1", "TestTag2")

	// Insert question with tags from tags slice, for checking tags questions
	_, err = questionService.CreateQuestion(&model.NewQuestionInput{
		Title:   "Test Question 1",
		Slug:    slug.Make("Test Question 1"),
		Author:  u.ID,
		Content: "Content of question one",
		Tags:    tags,
	})
	assert.NoError(t, err)

	// Check if page and itemsPerPage works
	itemsPerPageOne, _ := tagService.GetAllTags(1, 1)
	assert.Len(t, itemsPerPageOne, 1)
	itemsPerPageTwo, _ := tagService.GetAllTags(1, 2)
	assert.Len(t, itemsPerPageTwo, 2)
	secondPageResult, _ := tagService.GetAllTags(2, 1)
	assert.Len(t, secondPageResult, 1)

	// Check empty data result
	emptyDataResult, _ := tagService.GetAllTags(2, 4)
	assert.Len(t, emptyDataResult, 0)

	// Single Tag
	resultSingleTag, _ := tagService.GetTag(1)
	assert.Contains(t, resultSingleTag.Name, "TestTag1")

	// Return error if Tag does not exist
	_, err = tagService.GetTag(10)
	assert.Error(t, err)

	// Update Tag if exist
	_, err = tagService.UpdateTag(2, &model.UpdateTagInput{
		Name: "TestTag2Updated",
	})
	assert.NoError(t, err)

	// Check error if Tag does not exists
	_, err = tagService.UpdateTag(3, &model.UpdateTagInput{
		Name: "TestTag2Updated",
	})
	assert.Error(t, err)

	// Delete Tag
	err = tagService.DeleteTag(2)
	assert.NoError(t, err)

	// Check error if Tag does not exists
	err = tagService.DeleteTag(10)
	assert.Error(t, err)

	// Check if Tag is deleted
	deletedTagResult, _ := tagService.GetAllTags(1, 1)
	assert.Len(t, deletedTagResult, 1)

	// Checking tag questions
	tagQuestions, _ := tagService.GetTagQuestions(1)
	assert.Equal(t, tagQuestions.Edges.Questions[0].ID, 1)
	assert.Equal(t, tagQuestions.Edges.Questions[0].Title, "Test Question 1")
	assert.Len(t, tagQuestions.Edges.Questions, 1)

}
