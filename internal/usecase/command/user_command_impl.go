package command

import (
	"context"
	"errors"
	"ewallet-backend/internal/domain/entity"
	"ewallet-backend/internal/domain/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrMailAlreadyExists = errors.New("email already exists")
)

type userCommandHandler struct {
	cmdRepo   repository.UserCommandRepository
	queryRepo repository.UserQueryRepository
}

func NewUserCommandHandler(cmdRepo repository.UserCommandRepository, queryRepo repository.UserQueryRepository) UserCommandHandler {
	return &userCommandHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}
func (h *userCommandHandler) Register(ctx context.Context, input RegisterInput) (uuid.UUID, error) {
	exists, err := h.queryRepo.ExistsByEmail(ctx, input.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if exists {
		return uuid.Nil, ErrMailAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	user := entity.NewUser(input.FullName, input.Username, string(hashedPassword), input.Role)

	if err := h.cmdRepo.Create(ctx, user); err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
