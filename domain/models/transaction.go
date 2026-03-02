package models

import (
	"time"
)

type Transaction struct {
	ID                     uint   `gorm:"primaryKey;autoIncrement"`
	Amount                 Amount `gorm:"serializer:json"`
	ReferenceNumber        string `gorm:"varchar(255);not null"`
	MerchantID             string `gorm:"varchar(255);not null"`
	PartnerReferenceNumber string `gorm:"varchar(255);not null"`
	Status                 string `gorm:"varchar(25);not null"`
	QRContent              string `gorm:"text;not null"`
	TransactionDate        *time.Time
	PaidDate               *time.Time
}

type Amount struct {
	Value    uint   `json:"value"`
	Currency string `json:"currency"`
}
