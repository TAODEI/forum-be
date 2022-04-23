package token

import (
	"errors"
	"fmt"
	"time"

	"forum/log"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	jwtKey string

	ErrTokenInvalid = errors.New("The token is invalid.")
	ErrTokenExpired = errors.New("The token is expired.")
)

// getJwtKey get jwtKey.
func getJwtKey() string {
	if jwtKey == "" {
		jwtKey = viper.GetString("jwt_secret")
	}
	return jwtKey
}

// TokenPayload is a required payload when generates token.
type TokenPayload struct {
	ID      uint32        `json:"id"`
	Role    uint32        `json:"role"`
	TeamID  uint32        `json:"team_id"`
	Expired time.Duration `json:"expired"` // 有效时间（nanosecond）
}

// TokenResolve means returned payload when resolves token.
type TokenResolve struct {
	ID        uint32 `json:"id"`
	Role      uint32 `json:"role"`
	TeamID    uint32 `json:"team_id"`
	ExpiresAt int64  `json:"expires_at"` // 过期时间（时间戳，10位）
}

// GenerateToken generates token.
func GenerateToken(payload *TokenPayload) (string, error) {
	claims := &TokenClaims{
		ID:        payload.ID,
		Role:      payload.Role,
		TeamID:    payload.TeamID,
		ExpiresAt: time.Now().Unix() + int64(payload.Expired.Seconds()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(getJwtKey()))
}

// ResolveToken resolves token.
func ResolveToken(tokenStr string) (*TokenResolve, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(getJwtKey()), nil
	})

	if err != nil {
		log.Error("Token parsing failed because of an internal error", zap.String("cause", err.Error()))
		return nil, err
	}

	if !token.Valid {
		log.Error("Token parsing failed; the token is invalid.")
		return nil, ErrTokenInvalid
	}

	t := &TokenResolve{
		ID:        claims.ID,
		Role:      claims.Role,
		TeamID:    claims.TeamID,
		ExpiresAt: claims.ExpiresAt,
	}
	return t, nil
}
