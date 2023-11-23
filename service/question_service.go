package service

import (
	"context"
	"time"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/question"
	"github.com/sinisaos/chi-ent/model"

	"entgo.io/ent/dialect/sql"
	"github.com/gosimple/slug"
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

	s.Client.Question.UpdateOneID(question.ID).
		SetViews(question.Views + 1).
		Save(context.Background())

	return question, nil
}

func (s QuestionService) CreateQuestion(payload *model.NewQuestionInput) (*ent.Question, error) {
	var tags []*ent.Tag
	for _, v := range payload.Tags {
		tag, err := s.Client.Tag.Create().
			SetName(v).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	question, err := s.Client.Question.Create().
		SetTitle(payload.Title).
		SetSlug(slug.Make(payload.Content)).
		SetContent(payload.Content).
		SetCreatedAt(time.Now()).
		SetAuthorID(payload.Author).
		AddTags(tags...).
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

	var newTags []*ent.Tag
	for _, v := range payload.Tags {
		tag, err := s.Client.Tag.Create().
			SetName(v).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		newTags = append(newTags, tag)
	}

	question, err := existingQuestion.Update().
		SetTitle(payload.Title).
		SetSlug(slug.Make(payload.Content)).
		SetContent(payload.Content).
		RemoveTags(existingQuestion.Edges.Tags...).
		AddTags(newTags...).
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
		WithAuthor().
		WithAnswers().
		WithTags().
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
