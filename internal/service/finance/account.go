package finance

import (
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{
		accountRepository: accountRepository,
	}
}

func (as *AccountService) Create(account model.Account) (int64, error) {
	if err := as.accountRepository.DeactivateAllByAdm(account.AccountAdm); err != nil {
		return 0, err
	}

	account.AccountStatus = true

	return as.accountRepository.Insert(account)
}

func (as *AccountService) Update(account model.Account) error {
	if err := as.accountRepository.DeactivateAllByAdm(account.AccountAdm); err != nil {
		return err
	}

	account.AccountStatus = true

	return as.accountRepository.Update(account)
}

func (as *AccountService) Delete(accountID int64) error {
	// criar uma função para retornar true caso todas as contas de um adm estejam desativadas, forçando a ativar alguma conta

	return as.accountRepository.Delete(accountID)
}

func (as *AccountService) GetByID(accountID int64) (*model.Account, error) {
	return as.accountRepository.GetByID(accountID)
}

func (as *AccountService) Listing(admID int64) ([]model.AccountWithBank, error) {
	return as.accountRepository.Listing(admID)
}
