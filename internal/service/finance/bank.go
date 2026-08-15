package finance

import (
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type BankService struct {
	bankRepository *repository.BankRepository
}

func NewBankService(bankRepository *repository.BankRepository) *BankService {
	return &BankService{
		bankRepository: bankRepository,
	}
}

func (bs *BankService) Create(bank model.Bank) (int64, error) {
	return bs.bankRepository.Insert(bank)
}

func (bs *BankService) GetByID(bankID int64) (*model.Bank, error) {
	return bs.bankRepository.GetByID(bankID)
}

func (bs *BankService) Update(bank model.Bank) error {
	return bs.bankRepository.Update(bank)
}

func (bs *BankService) Delete(bankID int64) error {
	return bs.bankRepository.Delete(bankID)
}

func (bs *BankService) Listing() ([]model.Bank, error) {
	return bs.bankRepository.Listing()
}
