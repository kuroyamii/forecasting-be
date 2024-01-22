package authRepository

import (
	"context"
	"database/sql"
	"forecasting-be/internal/model"
	"forecasting-be/internal/query"

	"github.com/google/uuid"
)

type authRepository struct {
	DB *sql.DB
}

func NewAuthRepository(DB *sql.DB) authRepository {
	return authRepository{
		DB: DB,
	}
}

func (ar authRepository) CheckEmailAndUsername(ctx context.Context, username string, email string) (bool, error) {
	rows, err := ar.DB.QueryContext(ctx, query.GET_SAME_USERNAME_OR_EMAIL, username, email)
	if err != nil {
		return true, err
	}

	for rows.Next() {
		var count int
		err = rows.Scan(&count)
		if err != nil {
			return true, err
		}
		if count > 0 {
			return false, err
		}
	}
	return true, err
}

func (ar authRepository) BeginTrx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return ar.DB.BeginTx(ctx, opts)
}

func (ar authRepository) CommitTrx(trx *sql.Tx) error {
	return trx.Commit()
}

func (ar authRepository) RollbackTrx(trx *sql.Tx) error {
	return trx.Rollback()
}
func (ar authRepository) InsertUser(ctx context.Context, userRequest model.User) error {
	_, err := ar.DB.ExecContext(ctx, query.INSERT_USER,
		userRequest.ID,
		userRequest.Username,
		userRequest.Password,
		userRequest.Email,
		userRequest.FullName)
	return err
}

func (ar authRepository) SignIn(ctx context.Context, username string, password string) (model.User, error) {
	rows, err := ar.DB.QueryContext(ctx, query.GET_USER_BY_USERNAME_AND_PASSWORD, username, password)
	if err != nil {
		return model.User{}, err
	}
	var userData model.User
	for rows.Next() {
		err = rows.Scan(&userData.ID, &userData.Username, &userData.Email, &userData.FullName)
		if err != nil {
			return model.User{}, err
		}
	}
	return userData, err
}

func (ar authRepository) CheckUUID(ctx context.Context, id uuid.UUID) (bool, error) {

	return true, nil
}

func (ar authRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	rows, err := ar.DB.QueryContext(ctx, query.GET_USER_BY_ID, id)
	if err != nil {
		return model.User{}, err
	}
	var userData model.User
	for rows.Next() {
		err = rows.Scan(&userData.ID, &userData.Username, &userData.Email, &userData.FullName)
		if err != nil {
			return model.User{}, err
		}
	}
	return userData, err
}

// func (ar authRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
// 	_, err := ar.DB.ExecContext(ctx, query.DELETE_USER, userID)
// 	return err
// }
