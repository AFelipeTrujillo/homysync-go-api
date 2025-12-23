package ports

import "github.com/AFelipeTrujillo/homysync-go-api/internal/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
