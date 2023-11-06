package service

import (
	"context"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/user"
	"github.com/sinisaos/chi-ent/model"
)

type AuthService struct {
	Client *ent.Client
}

func NewAuthService(client *ent.Client) *AuthService {
	return &AuthService{
		Client: client,
	}
}

func (s AuthService) Login(payload *model.LoginUserInput) (*ent.User, error) {
	user, err := s.Client.User.Query().
		Where(user.Email(payload.Email)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return user, nil
}
