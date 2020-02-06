package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Invoice struct {
	gorm.Model
	InvoiceNumber string
	Date          time.Time
	SubscriptionID uint
}
