package repositories

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"manjo-test/domain/models"
)

type TransactionRepository struct {
	db *gorm.DB
}

type ITransactionRepository interface {
	FindByReferenceNumber(context.Context, string) (*models.Transaction, error)
	Create(context.Context, *models.Transaction) (*models.Transaction, error)
	Update(context.Context, string, *models.Transaction) (*models.Transaction, error)
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) FindByReferenceNumber(ctx context.Context, referenceNumber string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.WithContext(ctx).Where("reference_number = ?", referenceNumber).First(&transaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}

		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.WithContext(ctx).Create(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *TransactionRepository) Update(ctx context.Context, referenceNumber string, transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.WithContext(ctx).Where("reference_number = ?", referenceNumber).Updates(&transaction).Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
