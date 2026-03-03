package dto

import "time"

type CreateTransactionRequest struct {
	MerchantID             string `json:"merchantId" validate:"required"`
	Amount                 Amount `json:"amount" validate:"required"`
	PartnerReferenceNumber string `json:"partnerReferenceNo" validate:"required"`
}

type Amount struct {
	Value    string `json:"value" validate:"required"`
	Currency string `json:"currency" validate:"required"`
}

type CreateTransactionResponse struct {
	ReferenceNumber        string `json:"referenceNo"`
	PartnerReferenceNumber string `json:"partnerReferenceNo"`
	QRContent              string `json:"qrContent"`
}

type UpdateTransactionRequest struct {
	ReferenceNumber        string     `json:"originalReferenceNo" validate:"required"`
	PartnerReferenceNumber string     `json:"originalPartnerReferenceNo" validate:"required"`
	Status                 string     `json:"transactionStatusDesc" validate:"required"`
	PaidTime               *time.Time `json:"paidTime" validate:"required"`
	Amount                 Amount     `json:"amount" validate:"required"`
}

type UpdateTransactionResponse struct {
	ReferenceNumber string `json:"referenceNo"`
	Amount          Amount `json:"amount"`
	Status          string `json:"transactionStatusDesc"`
}

type TransactionResponse struct {
	MerchantID             string     `json:"merchantId"`
	ReferenceNumber        string     `json:"originalReferenceNo" `
	PartnerReferenceNumber string     `json:"originalPartnerReferenceNo" `
	Status                 string     `json:"transactionStatusDesc" `
	TransactionDate        *time.Time `json:"transactionDate" `
	PaidDate               *time.Time `json:"paidTime" `
	Amount                 Amount     `json:"amount"`
}
