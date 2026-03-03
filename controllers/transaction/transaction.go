package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"manjo-test/domain/dto"
	services "manjo-test/services/transaction"
)

type TransactionController struct {
	service services.ITransactionService
}

type ITransactionController interface {
	GetAll(*fiber.Ctx) error
	Create(*fiber.Ctx) error
	Update(*fiber.Ctx) error
}

func NewTransactionController(service services.ITransactionService) ITransactionController {
	return &TransactionController{service: service}
}

func (r *TransactionController) GetAll(ctx *fiber.Ctx) error {
	transactions, err := r.service.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    transactions,
	})
}

func (r *TransactionController) Create(ctx *fiber.Ctx) error {
	request := &dto.CreateTransactionRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	result, err := r.service.Create(ctx.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"responseCode":    "2004700",
		"responseMessage": "Successful",
		"referenceNo":     result.ReferenceNumber,
		"qrContent":       "00020101021226620015ID.CO.MANJO.WWW01189360085801751859910210EP278421820303UMI51530014ID.CO.QRIS.WWW0215ID102106515 192304121.0.21.09.255204481653033605502015802ID5904OLDI6013JAKARTA BARAT61051147062460525DIRECT-API-NMS-whhq7gvx5807031110806'ASPI663040FAD",
	})
}

func (r *TransactionController) Update(ctx *fiber.Ctx) error {
	request := &dto.UpdateTransactionRequest{}

	err := ctx.BodyParser(request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err = validate.Struct(request); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	result, err := r.service.Update(ctx.Context(), request.ReferenceNumber, request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"responseCode":          "2005100",
		"responseMessage":       "Successful",
		"transactionStatusDesc": result.Status,
	})
}
