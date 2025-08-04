package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Get(ctx context.Context, id uuid.UUID) (User, error) {
	user, err := gorm.G[User](u.db).Where("id = ?", id.String()).First(ctx)

	return user, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
