package info

import (
	"github.com/asvinicius/actnsgo/internal/client"
	"github.com/asvinicius/actnsgo/internal/response"
)

type StatusService struct {
	statusClient *client.Client
}

func NewStatusService(statutClient *client.Client) *StatusService {
	return &StatusService{
		statusClient: statutClient,
	}
}

func (s *StatusService) GetStatus() (*response.StatusResponse, error) {
	return s.statusClient.GetStatus()
}

/*

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

*/
