package jwt

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"github.com/Ndraaa15/Educode/pkg/errsx"
	"github.com/golang-jwt/jwt/v4"
)

func EncodeToken(user *entity.User) (string, error) {
	claims := &entity.JWTClaims{
		ID:   user.ID,
		Role: user.Role.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 3)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", errsx.NewCustomError(http.StatusInternalServerError, "Jwt.EncodeToken", "Failed to encode token", errors.New("ERROR_SIGNING_TOKEN"))
	}
	return signedToken, nil
}

func DecodeToken(token string) (*entity.JWTClaims, error) {
	decoded, err := jwt.ParseWithClaims(token, &entity.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if !decoded.Valid {
		return nil, errsx.NewCustomError(http.StatusInternalServerError, "Jwt.DecodeToken", "Failed to decode token", errors.New("INVALID_TOKEN"))
	}

	claims, ok := decoded.Claims.(*entity.JWTClaims)
	if !ok {
		return nil, errsx.NewCustomError(http.StatusInternalServerError, "Jwt,EncodeToken", "Failed to claim token", errors.New("ERROR_CLAIMING_TOKEN"))
	}

	return claims, nil
}
