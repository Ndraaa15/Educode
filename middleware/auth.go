package middleware

import (
	"net/http"
	"strings"

	"github.com/Ndraaa15/Educode/internal/pkg/entity"
	"github.com/Ndraaa15/Educode/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func ValidateJWTToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.Request.Header.Get("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You must be logged in first."})
			return
		}

		tokenParts := strings.SplitN(header, " ", 2)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token"})
			return
		}

		tokenString := tokenParts[1]

		claims, err := jwt.DecodeToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Failed Decode Token"})
			return
		}

		for _, role := range roles {
			if claims.Role == role {
				user := &entity.JWTClaims{
					ID:   claims.ID,
					Role: claims.Role,
				}

				ctx.Set("user", user)
				ctx.Next()
				return
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}
}
