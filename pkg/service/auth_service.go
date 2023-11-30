package service

import (
	"context"
	"time"

	"github.com/sinisaos/chi-ent/ent"
	"github.com/sinisaos/chi-ent/ent/user"
	"github.com/sinisaos/chi-ent/pkg/model"
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
	// update last login time
	s.Client.User.UpdateOneID(user.ID).
		SetLastLogin(time.Now()).
		Save(context.Background())

	return user, nil
}
