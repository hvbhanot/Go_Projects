package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecret"

// GenerateToken generates a JWT token for a given email and userId.
// It uses jwt.NewWithClaims to create a new token with the HMAC SHA256 signing method.
// The token contains the email, userId, and expiration time, which is set to 2 hours from the current time.
// The token is then signed with the secretKey and returned as a string.
// If any error occurs during token generation, an error is returned.
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// VerifyToken verifies the validity of the given token and extracts the userId.
// If the token is not parsed successfully, it returns an error with the message "could not parse token".
// If the token is invalid, it returns an error with the message "invalid Token".
// If the token claims are not valid, it returns an error with the message "invalid Token".
// Otherwise, it returns the extracted userId and nil error.
func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(Token *jwt.Token) (any, error) {
		_, ok := Token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	TokenIsvalid := parsedToken.Valid
	if !TokenIsvalid {
		return 0, errors.New("invalid Token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid Token")
	}
	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
