package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Subscription struct {
	gorm.Model
	TenantId  uint32
	StartDate time.Time
	EndDate   time.Time
	PlanID    uint
}
