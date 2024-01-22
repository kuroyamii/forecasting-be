package authService

import (
	"context"
	"forecasting-be/internal/dto"

	"github.com/google/uuid"
)

type AuthService interface {
	SignIn(ctx context.Context, username string, password string) (dto.SignInResponse, error)
	SignUp(ctx context.Context, userRequest dto.RegisterRequest) (uuid.UUID, error)
	RegenerateToken(ctx context.Context, rt string) (dto.SignInResponse, error)
	// CleanUpInvalidUser(ctx context.Context, userId uuid.UUID) error
}
