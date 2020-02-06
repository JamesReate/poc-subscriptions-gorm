package main

import (
	"fmt"
	"github.com/jamesreate/poc-subscriptions-gorm/internal/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=gorm dbname=subscriptions password=gorm sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.Subscription{}, &model.Invoice{}, &model.Plan{})

	// Create
	db.Create(&model.Plan{
		Name:  "Basic Plan",
	})
	db.Create(&model.Subscription{
		TenantId:  23,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Duration(100000)),
		PlanID:    1,
	})
	db.Create(&model.Invoice{
		InvoiceNumber:  "F009-2321321",
		Date:           time.Now(),
		SubscriptionID: 1,
	})
	var subscription model.Subscription
	db.First(&subscription, "tenant_id = ?", 23)

	fmt.Println(subscription)
	// Delete - delete product
	//db.Delete(&product)
}