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
	ar         authRepository.AuthRepository
	smtpConfig utilities.SMTPConfig
}

func NewAuthService(ar authRepository.AuthRepository, smtpConfig utilities.SMTPConfig) authService {
	return authService{
		ar:         ar,
		smtpConfig: smtpConfig,
	}
}

func (as authService) SignIn(ctx context.Context, username string, password string) (dto.SignInResponse, error) {
	// Hash password
	hashedPassword := utilities.HashPassword(password)

	// Check if the credential is valid
	userData, role, err := as.ar.SignIn(ctx, username, hashedPassword)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	// Insert queried user data to a struct without the ID
	data := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RoleID    int    `json:"role_id"`
	}{
		Username:  userData.Username,
		Email:     userData.Email,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		RoleID:    role,
	}

	// Creating a login token
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
	for {
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
	if !isAvailable {
		return uuid.UUID{}, errors.New("username or email is used")
	}

	// Check if the code is valid
	code, roleID, err := as.ar.GetCode(ctx, userRequest.AdminCode, userRequest.Email)
	if err != nil {
		return uuid.UUID{}, err
	}
	if code != userRequest.AdminCode {
		return uuid.UUID{}, errors.New("forbidden")
	}

	// Hash password
	password := utilities.HashPassword(userRequest.Password)
	// Assign values to model
	userRequestModel := model.User{
		ID:        id,
		Username:  userRequest.Username,
		Password:  password,
		Email:     userRequest.Email,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
	}

	// Insert to database
	err = as.ar.InsertUser(ctx, userRequestModel, roleID)
	if err != nil {
		trx.Rollback()
		return uuid.UUID{}, err
	}
	return uuid.UUID{}, trx.Commit()
}

func (as authService) RegenerateToken(ctx context.Context, rt string) (dto.SignInResponse, error) {
	// Get token env
	tokenEnv := utilities.GetTokenEnv()
	// Get data from refresh token
	id, created_at, err := utilities.GetDataFromRefreshToken(rt)
	if err != nil {
		return dto.SignInResponse{}, err
	}
	interval := time.Since(created_at)

	// If the refresh token is expired, return error. If not, proceed
	if interval.Hours() > float64(tokenEnv.RefreshTokenTTLHour) {
		return dto.SignInResponse{}, errors.New("refresh token invalid")
	}

	// Get user data with the user id stored in refresh token
	user, role, err := as.ar.GetUserByID(ctx, id)
	if err != nil {
		return dto.SignInResponse{}, err
	}

	// Initialize struct with queried user data
	data := struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		RoleID    int    `json:"role_id"`
	}{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		RoleID:    role,
	}
	// Create login tokens
	accessToken, refreshToken := utilities.CreateLoginToken(user.ID, data)
	res := dto.SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return res, nil
}

// This function is for inviting admin by adding code with corresponding email
func (as authService) AddAdmin(ctx context.Context, email string, role string) error {
	// Get role to be assigned
	trx, err := as.ar.BeginTrx(ctx, nil)
	if err != nil {
		return err
	}
	id, err := as.ar.GetRoleID(ctx, role)
	if err != nil {
		trx.Rollback()
		return err
	}
	// Create OTP code
	otp := utilities.GenerateOTP("6")

	// Insert email, role id, and OTP code
	err = as.ar.AddAdmin(ctx, email, id, otp)
	if err != nil {
		trx.Rollback()
		return err
	}
	// Send OTP to admin email
	err = utilities.SendOTPToEmail(otp, email)
	if err != nil {
		trx.Rollback()
		return err
	}
	return trx.Commit()
}
