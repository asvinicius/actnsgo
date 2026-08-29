package finance

import (
	"errors"
	"strconv"

	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
	accountService "github.com/asvinicius/actnsgo/internal/service/finance"
	"github.com/gofiber/fiber/v3"
)

type AccountHandler struct {
	accountService *accountService.AccountService
}

func NewAccountHandler(accountService *accountService.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (ah *AccountHandler) Create(c fiber.Ctx) error {
	accountBankStr := c.FormValue("account_bank")

	accountBank, err := strconv.ParseInt(accountBankStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_bank inválido",
		})
	}

	accountAdmStr := c.FormValue("account_adm")

	accountAdm, err := strconv.ParseInt(accountAdmStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_adm inválido",
		})
	}

	request := dto.AccountCreateRequest{
		AccountAdm:  accountAdm,
		AccountBank: accountBank,
		AccountKey:  c.FormValue("account_key"),
	}

	if request.AccountKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "chave pix é obrigatoria",
		})
	}

	account := model.Account{
		AccountAdm:    request.AccountAdm,
		AccountBank:   request.AccountBank,
		AccountKey:    request.AccountKey,
		AccountStatus: true,
	}

	accountID, err := ah.accountService.Create(account)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao cadastrar conta",
		})
	}

	account.AccountID = accountID

	return c.Status(fiber.StatusCreated).JSON(dto.ToAccountResponse(account))
}

func (ah *AccountHandler) Update(c fiber.Ctx) error {
	accountIDStr := c.Params("id")

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_id inválido",
		})
	}

	accountData, err := ah.accountService.GetByID(accountID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "conta não encontrada",
		})
	}

	accountBankStr := c.FormValue("account_bank")

	accountBank, err := strconv.ParseInt(accountBankStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_bank inválido",
		})
	}

	accountStatusStr := c.FormValue("account_status")

	accountStatus, err := strconv.ParseBool(accountStatusStr)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_status inválido",
		})
	}

	request := dto.AccountUpdateRequest{
		AccountBank:   accountBank,
		AccountKey:    c.FormValue("account_key"),
		AccountStatus: accountStatus,
	}

	if request.AccountKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "chave pix é obrigatoria",
		})
	}

	account := model.Account{
		AccountID:     accountData.AccountID,
		AccountAdm:    accountData.AccountAdm,
		AccountBank:   request.AccountBank,
		AccountKey:    request.AccountKey,
		AccountStatus: request.AccountStatus,
	}

	err = ah.accountService.Update(account)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao atualizar conta",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToAccountResponse(account))

}

func (ah *AccountHandler) Delete(c fiber.Ctx) error {
	accountIDStr := c.Params("id")

	accountID, err := strconv.ParseInt(accountIDStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_id inválido",
		})
	}

	err = ah.accountService.Delete(accountID)

	if errors.Is(err, repository.ErrAccountNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "conta não encontrada",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao remover conta",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "conta removida com sucesso",
	})
}

func (ah *AccountHandler) Listing(c fiber.Ctx) error {
	accountAdmStr := c.Params("id")

	accountAdm, err := strconv.ParseInt(accountAdmStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "account_adm inválido",
		})
	}

	listing, err := ah.accountService.Listing(accountAdm)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao listar contas",
		})
	}

	response := dto.ToAccountResponseList(listing)

	return c.Status(fiber.StatusOK).JSON(response)
}
