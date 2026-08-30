package adm

import (
	"fmt"

	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AdmService struct {
	admRepository *repository.AdmRepository
}

func NewAdmService(admRepository *repository.AdmRepository) *AdmService {
	return &AdmService{
		admRepository: admRepository,
	}
}

func (as *AdmService) Create(adm model.UserAdm) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(adm.AdmPassword), bcrypt.DefaultCost)

	if err != nil {
		return 0, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	adm.AdmPassword = string(hash)

	return as.admRepository.Insert(adm)
}

func (as *AdmService) Update(adm model.UserAdm) error {
	return as.admRepository.Update(adm)
}

func (as *AdmService) Delete(admID int64) error {
	return as.admRepository.Delete(admID)
}

func (as *AdmService) GetByID(admID int64) (*model.UserAdm, error) {
	return as.admRepository.GetByID(admID)
}

func (as *AdmService) Listing() ([]model.UserAdm, error) {
	return as.admRepository.Listing()
}
