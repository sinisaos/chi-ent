package service

import (
	"context"
	"time"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/user"
	"github.com/sinisaos/chi-ent/pkg/model"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Client *ent.Client
}

func NewUserService(client *ent.Client) *UserService {
	return &UserService{
		Client: client,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s UserService) GetAllUsers(page int, itemsPerPage int) ([]*ent.User, error) {
	offset := itemsPerPage * (page - 1)
	users, _ := s.Client.User.Query().
		Order(
			user.ByID(
				sql.OrderDesc(),
			),
		).
		WithQuestions().
		WithAnswers().
		WithTags().
		Limit(itemsPerPage).
		Offset(offset).
		All(context.Background())

	return users, nil
}

func (s UserService) GetUser(id int) (*ent.User, error) {
	user, err := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithAnswers().
		WithQuestions().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) CreateUser(payload *model.NewUserInput) (*ent.User, error) {
	hashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.Client.User.Create().
		SetEmail(payload.Email).
		SetUsername(payload.UserName).
		SetPassword(hashedPassword).
		SetCreatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) UpdateUser(id int, payload *model.UpdateUserInput) (*ent.User, error) {
	hashedPassword, err := hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	user, err := s.Client.User.UpdateOneID(id).
		SetUsername(payload.UserName).
		SetPassword(hashedPassword).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s UserService) DeleteUser(id int) error {
	err := s.Client.User.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) GetUserQuestions(id int) (*ent.User, error) {
	user, _ := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithQuestions().
		Only(context.Background())

	return user, nil
}

func (s UserService) GetUserAnswers(id int) (*ent.User, error) {
	user, _ := s.Client.User.Query().
		Where(user.IDEQ(id)).
		WithAnswers().
		Only(context.Background())

	return user, nil
}
