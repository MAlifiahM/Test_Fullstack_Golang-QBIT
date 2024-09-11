package auth

import (
	"case3/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type authService struct {
	authRepo domain.AuthRepository
	userRepo domain.UserRepository
	secret   string
}

func NewAuthService(authRepo domain.AuthRepository, userRepo domain.UserRepository, secret string) domain.AuthService {
	return &authService{
		authRepo: authRepo,
		userRepo: userRepo,
		secret:   secret,
	}
}

func (s *authService) Login(email string, password string) (map[string]string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	expiredTimeString := os.Getenv("EXPIRED_IN_HOURS")
	expiredTimeHrs, err := strconv.Atoi(expiredTimeString)

	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(time.Hour * time.Duration(expiredTimeHrs)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return nil, err
	}

	tokenData := map[string]string{"token": tokenString}

	return tokenData, nil
}

func (s *authService) Register(username string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	errs := s.authRepo.Register(&user)

	if errs != nil {
		return errs
	}

	return nil
}
