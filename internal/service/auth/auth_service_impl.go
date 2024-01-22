package authService

import (
	"context"
	"errors"
	"forecasting-be/internal/dto"
	"forecasting-be/internal/model"
	authRepository "forecasting-be/internal/repository/auth"
	"forecasting-be/pkg/utilities"
	"time"

	"github.com/google/uuid"
)

type authService struct {
	ar authRepository.AuthRepository
}

func NewAuthService(ar authRepository.AuthRepository) authService {
	return authService{
		ar: ar,
	}
}

func (as authService) SignIn(ctx context.Context, username string, password string) (dto.SignInResponse, error) {
	hashedPassword := utilities.HashPassword(password)
	userData, err := as.ar.SignIn(ctx, username, hashedPassword)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	data := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}{
		Username: userData.Username,
		Email:    userData.Email,
		FullName: userData.FullName,
	}

	accessToken, refreshToken := utilities.CreateLoginToken(userData.ID, data)
	return dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}
func (as authService) SignUp(ctx context.Context, userRequest dto.RegisterRequest) (uuid.UUID, error) {
	// Begin Transaction
	trx, err := as.ar.BeginTrx(ctx, nil)
	if err != nil {
		return uuid.UUID{}, err
	}

	// Checking if the UUID is available
	var id uuid.UUID
	for true {
		id = uuid.New()
		isAvailable, err := as.ar.CheckUUID(ctx, id)
		if err != nil {
			return uuid.UUID{}, err
		}
		if isAvailable {
			break
		}
	}
	// Check if the username or email isn't used
	isAvailable, err := as.ar.CheckEmailAndUsername(ctx, userRequest.Username, userRequest.Email)
	if err != nil {
		return uuid.UUID{}, err
	}
	if isAvailable == false {
		return uuid.UUID{}, errors.New("username or email is used!")
	}
	password := utilities.HashPassword(userRequest.Password)
	// Assign values to model
	userRequestModel := model.User{
		ID:       id,
		Username: userRequest.Username,
		Password: password,
		Email:    userRequest.Email,
		FullName: userRequest.FullName,
	}

	// Insert to database
	err = as.ar.InsertUser(ctx, userRequestModel)
	if err != nil {
		trx.Rollback()
		return uuid.UUID{}, err
	}
	return uuid.UUID{}, trx.Commit()
}

func (as authService) RegenerateToken(ctx context.Context, rt string) (dto.SignInResponse, error) {
	tokenEnv := utilities.GetTokenEnv()
	id, created_at, err := utilities.GetDataFromRefreshToken(rt)
	interval := time.Now().Sub(created_at)
	if interval.Hours() > float64(tokenEnv.RefreshTokenTTLHour) {
		return dto.SignInResponse{}, errors.New("refresh token invalid")
	}
	if err != nil {
		return dto.SignInResponse{}, err
	}
	user, err := as.ar.GetUserByID(ctx, id)
	data := struct {
		Username string
		Email    string
		FullName string
	}{}
	accessToken, refreshToken := utilities.CreateLoginToken(user.ID, data)
	res := dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}

// func (as authService) CleanUpInvalidUser(ctx context.Context, userId uuid.UUID) error {
// 	err := as.ar.DeleteUser(ctx, userId)
// 	return err
// }
