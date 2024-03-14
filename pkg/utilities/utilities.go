package utilities

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/xlzd/gotp"
)

var (
	Info  = Teal
	Warn  = Yellow
	Fatal = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

type TokenEnv struct {
	AccessTokenTTLHour  int64
	RefreshTokenTTLHour int64
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func JSONDecode(requestBody io.Reader, targetStruct interface{}) error {
	return json.NewDecoder(requestBody).Decode(targetStruct)
}

func HashPassword(password string) string {
	hash := sha256.New()
	modifiedPass := fmt.Sprint(password, os.Getenv("PASSWORD_HASH_KEY"))
	hash.Write([]byte(modifiedPass))
	passHash := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return passHash
}

func GetJWTKey() string {
	return os.Getenv("JWT_KEY")
}

func CreateLoginToken(userID uuid.UUID, data interface{}) (string, string) {
	claims := jwt.MapClaims{}
	claims["sub"] = userID
	claims["data"] = data
	claims["created_at"] = time.Now()
	tokenEnv := GetTokenEnv()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenEnv.AccessTokenTTLHour)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte(GetJWTKey()))

	claims = jwt.MapClaims{}
	claims["sub"] = userID
	claims["created_at"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenEnv.RefreshTokenTTLHour)).Unix()
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, _ := rtoken.SignedString([]byte(GetJWTKey()))

	return accessToken, refreshToken
}

func GetTokenEnv() TokenEnv {
	tokenTTL, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_TTL_HOUR"), 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	refreshTTL, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_TTL_HOUR"), 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	return TokenEnv{
		AccessTokenTTLHour:  tokenTTL,
		RefreshTokenTTLHour: refreshTTL,
	}
}

func GetDataFromRefreshToken(rt string) (uuid.UUID, time.Time, error) {
	jwtKey := GetJWTKey()
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(rt, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}
	if token.Valid {
		var res time.Time
		switch exp := claims["exp"].(type) {
		case float64:
			res = time.Unix(int64(exp), 0)
		case json.Number:
			v, _ := exp.Int64()
			res = time.Unix(v, 0)
		}
		if res.Before(time.Now()) {
			return uuid.UUID{}, time.Time{}, errors.New("token expired")
		}
		idStr := fmt.Sprintf("%v", claims["sub"])
		idConv, err := uuid.Parse(idStr)
		if err != nil {
			return uuid.UUID{}, time.Time{}, err
		}
		return idConv, res, nil
	}
	return uuid.UUID{}, time.Time{}, errors.New("token invalid")
}

func GenerateOTP(otpLengthStr string) string {
	// otpLengthStr := os.Getenv("OTP_LENGTH")
	otpLength, err := strconv.ParseInt(otpLengthStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	length := int(otpLength)
	str := gotp.RandomSecret(length)
	str = str[:otpLength]
	return str
}
