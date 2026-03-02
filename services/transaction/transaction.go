package services

import (
	"context"
	"errors"
	"fmt"
	"manjo-test/domain/dto"
	"manjo-test/domain/models"
	repositories "manjo-test/repositories/transaction"
	"strconv"
	"time"
)

type TransactionService struct {
	repository repositories.ITransactionRepository
}

type ITransactionService interface {
	Create(context.Context, *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error)
	Update(context.Context, string, *dto.UpdateTransactionRequest) (*dto.UpdateTransactionResponse, error)
}

func NewTransactionService(repository repositories.ITransactionRepository) ITransactionService {
	return &TransactionService{repository: repository}
}

func (r *TransactionService) Create(ctx context.Context, request *dto.CreateTransactionRequest) (*dto.CreateTransactionResponse, error) {
	value, err := strconv.ParseFloat(request.Amount.Value, 64)
	if err != nil {
		return nil, errors.New("invalid amount value")
	}

	if value <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	referenceNumber := fmt.Sprintf("A%010d", time.Now().UnixNano()%10000000000)

	now := time.Now()

	transaction := &models.Transaction{
		PartnerReferenceNumber: request.PartnerReferenceNumber,
		Amount: models.Amount{
			Value:    uint(value),
			Currency: request.Amount.Currency,
		},
		ReferenceNumber: referenceNumber,
		TransactionDate: &now,
	}

	newTransaction, err := r.repository.Create(ctx, transaction)
	if err != nil {
		return nil, err
	}

	transactionResult := &dto.CreateTransactionResponse{
		ReferenceNumber:        newTransaction.ReferenceNumber,
		PartnerReferenceNumber: newTransaction.PartnerReferenceNumber,
		QRContent:              newTransaction.QRContent,
	}

	return transactionResult, nil
}

func (r *TransactionService) Update(ctx context.Context, referenceNumber string, request *dto.UpdateTransactionRequest) (*dto.UpdateTransactionResponse, error) {
	_, err := r.repository.FindByReferenceNumber(ctx, referenceNumber)
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseFloat(request.Amount.Value, 64)
	if err != nil {
		return nil, errors.New("invalid amount value")
	}

	if value <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	updateTransaction := &models.Transaction{
		PartnerReferenceNumber: request.PartnerReferenceNumber,
		Amount: models.Amount{
			Value:    uint(value),
			Currency: request.Amount.Currency,
		},
		Status:   "Success",
		PaidDate: request.PaidTime,
	}

	result, err := r.repository.Update(ctx, referenceNumber, updateTransaction)
	if err != nil {
		return nil, err
	}

	transactionResponse := &dto.UpdateTransactionResponse{
		ReferenceNumber: result.ReferenceNumber,
		Amount: dto.Amount{
			Value:    strconv.Itoa(int(result.Amount.Value)),
			Currency: result.Amount.Currency,
		},
		Status: result.Status,
	}

	return transactionResponse, nil
}
