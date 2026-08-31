package adm

import (
	"errors"
	"strconv"
	"time"

	"github.com/asvinicius/actnsgo/internal/dto"
	"github.com/asvinicius/actnsgo/internal/model"
	"github.com/asvinicius/actnsgo/internal/repository"
	admservice "github.com/asvinicius/actnsgo/internal/service/adm"
	"github.com/gofiber/fiber/v3"
)

type AdmHandler struct {
	admService *admservice.AdmService
}

func NewAdmHandler(admService *admservice.AdmService) *AdmHandler {
	return &AdmHandler{
		admService: admService,
	}
}

func (ah *AdmHandler) Create(c fiber.Ctx) error {
	var request dto.AdmRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dados inválidos",
		})
	}

	if request.AdmName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "nome do administrador é obrigatorio",
		})
	}

	if request.AdmLogin == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "login do administrador é obrigatorio",
		})
	}

	adm := model.UserAdm{
		AdmName:      request.AdmName,
		AdmLogin:     request.AdmLogin,
		AdmPassword:  request.AdmLogin,
		AdmStatus:    true,
		AdmCreatedAt: time.Now(),
	}

	admID, err := ah.admService.Create(adm)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao cadastrar administrador",
		})
	}

	adm.AdmID = admID

	return c.Status(fiber.StatusCreated).JSON(dto.ToAdmResponse(adm))
}

func (ah *AdmHandler) Update(c fiber.Ctx) error {
	var request dto.AdmUpdateRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dados inválidos",
		})
	}

	if request.AdmName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "nome do administrador é obrigatorio",
		})
	}

	if request.AdmLogin == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "login do administrador é obrigatorio",
		})
	}

	admIDStr := c.Params("id")

	admID, err := strconv.ParseInt(admIDStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "adm_id inválido",
		})
	}

	admData, err := ah.admService.GetByID(admID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "administrador não encontrado",
		})
	}

	updatedAt := time.Now()

	adm := model.UserAdm{
		AdmID:        admData.AdmID,
		AdmName:      request.AdmName,
		AdmLogin:     request.AdmLogin,
		AdmPassword:  admData.AdmPassword,
		AdmStatus:    request.AdmStatus,
		AdmCreatedAt: admData.AdmCreatedAt,
		AdmUpdatedAt: &updatedAt,
		AdmLastLogin: admData.AdmLastLogin,
	}

	err = ah.admService.Update(adm)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao atualizar administrador",
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToAdmResponse(adm))
}

func (ah *AdmHandler) Delete(c fiber.Ctx) error {
	admIDStr := c.Params("id")

	admID, err := strconv.ParseInt(admIDStr, 10, 64)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "adm_id inválido",
		})
	}

	err = ah.admService.Delete(admID)

	if errors.Is(err, repository.ErrAdmNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "administrador não encontrado",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao remover administrador",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "administrador removido com sucesso",
	})
}

func (ah *AdmHandler) Listing(c fiber.Ctx) error {
	listing, err := ah.admService.Listing()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "erro ao listar administradores",
		})
	}

	response := dto.ToAdmResponseList(listing)

	return c.Status(fiber.StatusOK).JSON(response)
}
