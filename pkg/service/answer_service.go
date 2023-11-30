package service

import (
	"context"
	"time"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/answer"
	"github.com/sinisaos/chi-ent/pkg/model"

	"entgo.io/ent/dialect/sql"
)

type AnswerService struct {
	Client *ent.Client
}

func NewAnswerService(client *ent.Client) *AnswerService {
	return &AnswerService{
		Client: client,
	}
}

func (s AnswerService) GetAllAnswers(page int, itemsPerPage int) ([]*ent.Answer, error) {
	offset := itemsPerPage * (page - 1)
	answers, _ := s.Client.Answer.Query().
		WithAuthor().
		WithQuestion().
		Order(
			answer.ByID(
				sql.OrderDesc(),
			),
		).
		Limit(itemsPerPage).
		Offset(offset).
		All(context.Background())

	return answers, nil
}

func (s AnswerService) GetAnswer(id int) (*ent.Answer, error) {
	answer, err := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithAuthor().
		WithQuestion().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) CreateAnswer(payload *model.NewAnswerInput) (*ent.Answer, error) {
	answer, err := s.Client.Answer.Create().
		SetContent(payload.Content).
		SetCreatedAt(time.Now()).
		SetAuthorID(payload.Author).
		SetQuestionID(payload.Question).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) UpdateAnswer(id int, payload *model.UpdateAnswerInput) (*ent.Answer, error) {
	answer, err := s.Client.Answer.UpdateOneID(id).
		SetContent(payload.Content).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (s AnswerService) DeleteAnswer(id int) error {
	err := s.Client.Answer.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s AnswerService) GetAnswerAuthor(id int) (*ent.Answer, error) {
	answer, _ := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithAuthor().
		Only(context.Background())

	return answer, nil
}

func (s AnswerService) GetAnswerQuestion(id int) (*ent.Answer, error) {
	answer, _ := s.Client.Answer.Query().
		Where(answer.IDEQ(id)).
		WithQuestion().
		Only(context.Background())

	return answer, nil
}
