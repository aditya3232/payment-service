package models

import "time"

type Payment struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	InvoiceID   int       `gorm:"not null;index"`
	Amount      float64   `gorm:"not null"`
	Method      string    `gorm:"type:varchar(50);not null"`
	ReferenceNo string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	PaidAt      time.Time `gorm:"not null"`
	CreatedAt   *time.Time
}
