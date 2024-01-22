package authRepository

import (
	"context"
	"database/sql"
	"forecasting-be/internal/model"

	"github.com/google/uuid"
)

type AuthRepository interface {
	CheckEmailAndUsername(ctx context.Context, username string, email string) (bool, error)
	InsertUser(ctx context.Context, userRequest model.User) error
	SignIn(ctx context.Context, username string, password string) (model.User, error)
	CheckUUID(ctx context.Context, id uuid.UUID) (bool, error)
	BeginTrx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	CommitTrx(trx *sql.Tx) error
	RollbackTrx(trx *sql.Tx) error
	GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error)
	// DeleteUser(ctx context.Context, userID uuid.UUID) error
}
