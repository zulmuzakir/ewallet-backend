package repository

import (
	"context"
	"ewallet-backend/internal/domain/entity"
	"ewallet-backend/internal/domain/repository"
	"ewallet-backend/internal/repository/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userCommandRepo struct {
	queries *sqlc.Queries
}

func NewUserCommandRepository(db *pgxpool.Pool) repository.UserCommandRepository {
	return &userCommandRepo{
		queries: sqlc.New(db),
	}
}

func (r *userCommandRepo) Create(ctx context.Context, user *entity.User) error {
	now := time.Now()

	_, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:        user.ID,
		FullName:  user.FullName,
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: now,
		UpdatedAt: now,
	})

	return err
}
