package repository

import (
	"github.com/AFelipeTrujillo/homysync-go-api/internal/domain"
	"gorm.io/gorm"
)

// GormUserRepository is a GORM-based implementation of the UserRepository interface.
type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// FindByEmail retrieves a user by their email address.
// (r *GormUserRepository) is the receiver for the method, indicating it belongs to GormUserRepository.
func (r *GormUserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
