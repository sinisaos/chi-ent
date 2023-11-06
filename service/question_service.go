package service

import (
	"context"
	"time"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/model"

	"entgo.io/ent/dialect/sql"
)

type QuestionService struct {
	Client *ent.Client
}

func NewQuestionService(client *ent.Client) *QuestionService {
	return &QuestionService{
		Client: client,
	}
}

func (s QuestionService) GetAllQuestions(page int, itemsPerPage int) ([]*ent.Question, error) {
	offset := itemsPerPage * (page - 1)
	questions, _ := s.Client.Question.Query().
		WithAuthor().
		WithTags().
		Order(
			question.ByID(
				sql.OrderDesc(),
			),
		).
		Limit(itemsPerPage).
		Offset(offset).
		All(context.Background())

	return questions, nil
}

func (s QuestionService) GetQuestion(id int) (*ent.Question, error) {
	question, err := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAuthor().
		WithAnswers().
		WithTags().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) CreateQuestion(payload *model.NewQuestionInput) (*ent.Question, error) {
	question, err := s.Client.Question.Create().
		SetTitle(payload.Title).
		SetContent(payload.Content).
		SetCreatedAt(time.Now()).
		SetAuthorID(payload.Author).
		AddTagIDs(payload.Tags...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) UpdateQuestion(id int, payload *model.UpdateQuestionInput) (*ent.Question, error) {
	existingQuestion, err := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithTags().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	var tagsSlice []int

	for i := range existingQuestion.Edges.Tags {
		tagsSlice = append(tagsSlice, existingQuestion.Edges.Tags[i].ID)
	}

	question, err := existingQuestion.Update().
		SetTitle(payload.Title).
		SetContent(payload.Content).
		RemoveTagIDs(tagsSlice...).
		AddTagIDs(payload.Tags...).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (s QuestionService) DeleteQuestion(id int) error {
	err := s.Client.Question.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s QuestionService) GetQuestionAnswers(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAnswers().
		Only(context.Background())

	return question, nil
}

func (s QuestionService) GetQuestionAuthor(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithAuthor().
		Only(context.Background())

	return question, nil
}

func (s QuestionService) GetQuestionTags(id int) (*ent.Question, error) {
	question, _ := s.Client.Question.Query().
		Where(question.IDEQ(id)).
		WithTags().
		Only(context.Background())

	return question, nil
}
