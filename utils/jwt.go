package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// SecretKey is the secret used for signing JWT tokens
var SecretKey = []byte("secret_key") // Replace with a strong secret key

// GenerateToken generates a JWT token for a given user ID
func GenerateToken(userID string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	// Create the claims
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the user ID if valid
func ValidateToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return SecretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Extract claims
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", errors.New("invalid token")
}