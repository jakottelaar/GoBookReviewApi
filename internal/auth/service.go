package auth

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type AuthService interface {
	Login(req *LoginRequest) (string, error)
	Register(req *RegisterRequest) (*user.User, error)
}

type authService struct {
	repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

var tokenAuth *jwtauth.JWTAuth

func (s *authService) Login(req *LoginRequest) (string, error) {

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	_, err = argon2id.ComparePasswordAndHash(req.Password, user.Password)
	if err != nil {
		return "", err
	}

	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil, jwt.WithAcceptableSkew(30*time.Second))

	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": user.ID}) //Handle err or security risk?

	return tokenString, nil

}

func (s *authService) Register(userReq *RegisterRequest) (*user.User, error) {

	newId := uuid.New()

	// Add password hashing here
	hashedPassword, err := argon2id.CreateHash(userReq.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		ID:       newId,
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: hashedPassword,
	}

	savedUser, err := s.repo.Save(newUser)

	if err != nil {
		return nil, err
	}

	return savedUser, nil

}
