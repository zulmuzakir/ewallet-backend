package repository

import "context"

type UserQueryRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}