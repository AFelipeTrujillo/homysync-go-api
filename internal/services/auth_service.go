package services

import (
	"github.com/AFelipeTrujillo/homysync-go-api/internal/core/ports"
	"github.com/AFelipeTrujillo/homysync-go-api/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo ports.UserRepository
}

func NewAuthService(repo ports.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(email, password, name string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		GlobalName:   name,
	}

	return s.repo.Create(user)
}
