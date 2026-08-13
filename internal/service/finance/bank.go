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

func (bs *BankService) Listing() ([]model.Bank, error) {
	return bs.bankRepository.Listing()
}
