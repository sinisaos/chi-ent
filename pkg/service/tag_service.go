package service

import (
	"context"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/tag"
	"github.com/sinisaos/chi-ent/pkg/model"

	"entgo.io/ent/dialect/sql"
)

type TagService struct {
	Client *ent.Client
}

func NewTagService(client *ent.Client) *TagService {
	return &TagService{
		Client: client,
	}
}

func (s TagService) GetAllTags(page int, itemsPerPage int) ([]*ent.Tag, error) {
	offset := itemsPerPage * (page - 1)
	tags, _ := s.Client.Tag.Query().
		WithQuestions().
		Order(
			tag.ByID(
				sql.OrderDesc(),
			),
		).
		Limit(itemsPerPage).
		Offset(offset).
		All(context.Background())

	return tags, nil
}

func (s TagService) GetTag(id int) (*ent.Tag, error) {
	tag, err := s.Client.Tag.Query().
		Where(tag.IDEQ(id)).
		WithQuestions().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagService) UpdateTag(id int, payload *model.UpdateTagInput) (*ent.Tag, error) {
	tag, err := s.Client.Tag.UpdateOneID(id).
		SetName(payload.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s TagService) DeleteTag(id int) error {
	err := s.Client.Tag.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s TagService) GetTagQuestions(id int) (*ent.Tag, error) {
	tag, _ := s.Client.Tag.Query().
		Where(tag.IDEQ(id)).
		WithQuestions().
		Only(context.Background())

	return tag, nil
}
