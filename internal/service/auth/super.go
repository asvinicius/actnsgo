package auth

import (
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
	"github.com/asvinicius/actnsgo/internal/service/token"
	"golang.org/x/crypto/bcrypt"
)

var ErrSuperNotActive = errors.New("super not active")
var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	superRepository *repository.SuperRepository
	tokenService    *token.TokenService
}

func NewAuthService(superRepository *repository.SuperRepository, tokenService *token.TokenService) *AuthService {
	return &AuthService{
		superRepository: superRepository,
		tokenService:    tokenService,
	}
}

func (s *AuthService) Authenticate(superLogin, superPassword string) (*model.UserSuper, string, error) {

	super, err := s.superRepository.FindByLogin(superLogin)

	if !super.SuperStatus {
		return nil, "", ErrSuperNotActive
	}

	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(super.SuperPassword), []byte(superPassword)); err != nil {
		return nil, "", ErrInvalidCredentials
	}

	token, err := s.tokenService.GenerateSuperToken(super)

	if err != nil {
		return nil, "", err
	}

	return super, token, nil
}
