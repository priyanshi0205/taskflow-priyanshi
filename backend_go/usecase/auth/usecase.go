package auth

import (
	"net/http"
	"strings"
	"time"

	"backedn_go/internal/shared/apperror"
	authrepo "backedn_go/repository/auth"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type usecase struct {
	repo      authrepo.Repository
	jwtSecret string
}

func New(repo authrepo.Repository, jwtSecret string) Usecase {
	return &usecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *usecase) Register(input RegisterInput) (*RegisterOutput, *apperror.AppError) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "failed to process password")
	}

	user, err := u.repo.CreateUser(input.Name, input.Email, string(hashedPassword))
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return nil, apperror.New(http.StatusBadRequest, "Email already exists")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	return &RegisterOutput{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (u *usecase) Login(input LoginInput) (*LoginOutput, *apperror.AppError) {
	user, err := u.repo.FindUserByEmail(input.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.New(http.StatusUnauthorized, "Invalid credentials")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, apperror.New(http.StatusUnauthorized, "Invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	return &LoginOutput{Token: tokenString}, nil
}
