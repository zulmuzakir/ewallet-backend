package command

import (
	"context"

	"github.com/google/uuid"
)

type RegisterInput struct {
	FullName string
	Username string
	Email    string
	Password string
	Role     string
}

type UserCommandHandler interface {
	Register(ctx context.Context, input RegisterInput) (uuid.UUID, error)
}
