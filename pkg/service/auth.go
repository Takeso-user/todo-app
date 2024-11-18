package service

import (
	"crypto/sha1"
	"fmt"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/Takeso-user/todo-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "asd235kjkbhjbk988n(8TTRasd"
	signingKey = "asd235kjkbhjbk988##28TTRasd"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todoapp.User) (int, error) {
	user.Password, _ = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	p, _ := GeneratePasswordHash(password)
	user, err := s.repo.GetUser(username, p)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.ID,
		})
	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) ParseToken(token string) (int, error) {
	tkn, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	if !tkn.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	claims, ok := tkn.Claims.(*tokenClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims")
	}
	return claims.UserId, nil
}

func GeneratePasswordHash(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(salt))), nil
}
