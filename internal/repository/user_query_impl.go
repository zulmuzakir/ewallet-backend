package repository

import (
	"context"
	"ewallet-backend/internal/domain/repository"
	"ewallet-backend/internal/repository/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userQueryRepo struct {
	queries *sqlc.Queries
}

func NewUserQueryRepository(db *pgxpool.Pool) repository.UserQueryRepository {
	return &userQueryRepo{
		queries: sqlc.New(db),
	}
}

func (r *userQueryRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.queries.ExistsUserByEmail(ctx, email)
}