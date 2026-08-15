package finance

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
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
		publicPath := "/uploads/banks/" + fileName

		if err := c.SaveFile(file, fullPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "erro ao salvar imagem",
			})
		}

		logoPath = &publicPath
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

func (bh *BankHandler) Update(c fiber.Ctx) error {
	bankIDStr := c.Params("id")

	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bank_id inválido",
		})
	}

	bankData, err := bh.bankService.GetByID(bankID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "banco não encontrado",
		})
	}

	request := dto.BankRequest{
		BankName: c.FormValue("edit_bank_name"),
	}

	if request.BankName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "nome do banco é obrigatorio",
		})
	}

	logoPath := bankData.BankLogo

	file, err := c.FormFile("edit_bank_logo")

	if err == nil {
		if bankData.BankLogo != nil {
			oldFileName := filepath.Base(*bankData.BankLogo)
			oldFullPath := filepath.Join(bh.uploadDir, oldFileName)
			os.Remove(oldFullPath)
		}

		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("bank_%d%s", time.Now().UnixNano(), ext)
		fullPath := filepath.Join(bh.uploadDir, fileName)
		publicPath := "/uploads/banks/" + fileName

		if err := c.SaveFile(file, fullPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "erro ao salvar imagem",
			})
		}

		logoPath = &publicPath
	}

	bank := model.Bank{
		BankID:     bankID,
		BankName:   request.BankName,
		BankLogo:   logoPath,
		BankStatus: bankData.BankStatus,
	}

	err = bh.bankService.Update(bank)

	if errors.Is(err, repository.ErrBankNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "banco não encontrado",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao atualizar banco",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToBankResponse(bank))
}

func (bh *BankHandler) Delete(c fiber.Ctx) error {
	bankIDStr := c.Params("id")

	bankID, err := strconv.ParseInt(bankIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "bank_id inválido",
		})
	}

	bankData, err := bh.bankService.GetByID(bankID)

	if err == nil {
		if bankData.BankLogo != nil {
			oldFileName := filepath.Base(*bankData.BankLogo)
			oldFullPath := filepath.Join(bh.uploadDir, oldFileName)
			os.Remove(oldFullPath)
		}
	}

	err = bh.bankService.Delete(bankID)

	if errors.Is(err, repository.ErrBankNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "banco não encontrado",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao remover banco",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "banco removido com sucesso",
	})
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
