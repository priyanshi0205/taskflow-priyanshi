package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const ClaimsContextKey = "auth_claims"

type AuthClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			slog.Warn(
				"authentication failed: missing bearer token",
				"http_method", c.Request.Method,
				"http_path", c.Request.URL.Path,
				"http_route", c.FullPath(),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &AuthClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			slog.Warn(
				"authentication failed: invalid token",
				"http_method", c.Request.Method,
				"http_path", c.Request.URL.Path,
				"http_route", c.FullPath(),
				"error", err,
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		if strings.TrimSpace(claims.UserID) == "" {
			slog.Warn(
				"authentication failed: token missing user_id",
				"http_method", c.Request.Method,
				"http_path", c.Request.URL.Path,
				"http_route", c.FullPath(),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set(ClaimsContextKey, claims)
		c.Next()
	}
}

func ClaimsFromContext(c *gin.Context) (*AuthClaims, bool) {
	value, ok := c.Get(ClaimsContextKey)
	if !ok {
		return nil, false
	}
	claims, ok := value.(*AuthClaims)
	return claims, ok
}

func GetUserIDfromContext(c *gin.Context) string {
	return UserIDFromContext(c)
}

func UserIDFromContext(c *gin.Context) string {
	claims, ok := ClaimsFromContext(c)
	if !ok || claims == nil {
		return ""
	}
	return strings.TrimSpace(claims.UserID)
}
