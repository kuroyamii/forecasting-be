package authRepository

import (
	"context"
	"database/sql"
	"errors"
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

func (ar authRepository) GetRoleID(ctx context.Context, role string) (int, error) {
	// Select roles with role name
	rows, err := ar.DB.QueryContext(ctx, query.GET_ROLE_ID, role)
	if err != nil {
		return 0, err
	}
	// Init integer variable
	var id int
	if rows.Next() {
		// Assign value
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	return 0, errors.New("something went wrong")
}
func (ar authRepository) AddAdmin(ctx context.Context, email string, role int, code string) error {
	// Execute insert script
	_, err := ar.DB.ExecContext(ctx, query.INVITE_ADMIN, email, role, code)
	return err
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
func (ar authRepository) InsertUser(ctx context.Context, userRequest model.User, roleID int) error {
	// Execute insert script
	_, err := ar.DB.ExecContext(ctx, query.INSERT_USER,
		userRequest.ID,
		userRequest.Username,
		userRequest.Password,
		userRequest.Email,
		userRequest.FirstName,
		userRequest.LastName,
		roleID)
	return err
}

func (ar authRepository) SignIn(ctx context.Context, username string, password string) (model.User, int, error) {
	// Select corresponding user
	rows, err := ar.DB.QueryContext(ctx, query.GET_USER_BY_USERNAME_AND_PASSWORD, username, password)
	if err != nil {
		return model.User{}, 0, err
	}
	// Initialize empty user
	var userData model.User
	var roleID int

	for rows.Next() {
		// Assign value
		err = rows.Scan(&userData.ID, &userData.Username, &userData.Email, &userData.FirstName, &userData.LastName, &roleID)
		if err != nil {
			return model.User{}, 0, err
		}
	}
	if roleID == 0 {
		return model.User{}, 0, errors.New("unauthorized")
	}
	return userData, roleID, err
}

func (ar authRepository) CheckUUID(ctx context.Context, id uuid.UUID) (bool, error) {

	return true, nil
}

func (ar authRepository) GetUserByID(ctx context.Context, id uuid.UUID) (model.User, int, error) {
	// Select user data by ID
	rows, err := ar.DB.QueryContext(ctx, query.GET_USER_BY_ID, id)
	if err != nil {
		return model.User{}, 0, err
	}
	// Initialize empty variable
	var userData model.User
	var roleID int
	for rows.Next() {
		// Assign data
		err = rows.Scan(&userData.ID, &userData.Username, &userData.Email, &userData.FirstName, &userData.LastName, &roleID)
		if err != nil {
			return model.User{}, 0, err
		}
	}
	return userData, roleID, err
}

func (ar authRepository) GetCode(ctx context.Context, code string, email string) (string, int, error) {
	rows, err := ar.DB.QueryContext(ctx, query.GET_CODE, code, email)
	if err != nil {
		return "", 0, err
	}
	var result string
	var roleID int

	for rows.Next() {
		err = rows.Scan(&result, &roleID)
		if err != nil {
			return "", 0, err
		}
	}
	return result, roleID, nil
}

// func (ar authRepository) DeleteUser(ctx context.Context, userID uuid.UUID) error {
// 	_, err := ar.DB.ExecContext(ctx, query.DELETE_USER, userID)
// 	return err
// }
