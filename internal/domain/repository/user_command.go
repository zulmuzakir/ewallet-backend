package repository

import (
	"context"
	"ewallet-backend/internal/domain/entity"
)

type UserCommandRepository interface {
	Create(ctx context.Context, user *entity.User) error
}
