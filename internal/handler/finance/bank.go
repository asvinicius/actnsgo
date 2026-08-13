package finance

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/model"
	financeservice "github.com/asvinicius/actnsgo/internal/service/finance"
	"github.com/gofiber/fiber/v3"
)

type BankHandler struct {
	bankService *financeservice.BankService
	uploadDir   string
}

func NewBankHandler(bankService *financeservice.BankService, uploadDir string) *BankHandler {
	return &BankHandler{
		bankService: bankService,
		uploadDir:   uploadDir,
	}
}

func (bh *BankHandler) Create(c fiber.Ctx) error {
	request := dto.BankRequest{
		BankName: c.FormValue("bank_name"),
	}

	if request.BankName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "nome do banco é obrigatorio",
		})
	}

	var logoPath *string

	file, err := c.FormFile("bank_logo")

	if err == nil {
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("bank_%d%s", time.Now().UnixNano(), ext)
		fullPath := filepath.Join(bh.uploadDir, fileName)

		if err := c.SaveFile(file, fullPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "erro ao salvar imagem",
			})
		}

		logoPath = &fullPath
	}

	bank := model.Bank{
		BankName:   request.BankName,
		BankLogo:   logoPath,
		BankStatus: true,
	}

	bankID, err := bh.bankService.Create(bank)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao cadastrar banco",
		})
	}

	bank.BankID = bankID

	return c.Status(fiber.StatusCreated).JSON(dto.ToBankResponse(bank))
}

func (bh *BankHandler) Listing(c fiber.Ctx) error {
	listing, err := bh.bankService.Listing()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao listar bancos",
		})
	}

	response := dto.ToBankResponseList(listing)

	return c.Status(fiber.StatusOK).JSON(response)
}
