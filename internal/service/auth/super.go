package auth

import (
	"errors"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrSuperNotActive = errors.New("super not active")
var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService struct {
	superRepository *repository.SuperRepository
}

func NewAuthService(superRepository *repository.SuperRepository) *AuthService {
	return &AuthService{
		superRepository: superRepository,
	}
}

func (s *AuthService) Authenticate(superLogin, superPassword string) (*model.UserSuper, error) {

	super, err := s.superRepository.FindByLogin(superLogin)

	if !super.SuperStatus {
		return nil, ErrSuperNotActive
	}

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(super.SuperPassword), []byte(superPassword)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return super, nil
}
