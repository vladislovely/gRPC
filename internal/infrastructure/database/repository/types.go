package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	Get(ctx context.Context, id uuid.UUID) (User, error)
}

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}
