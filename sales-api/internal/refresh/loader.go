package refresh

import (
	"encoding/csv"
	"log"
	"os"
	"sales-api/internal/db"
	"sales-api/internal/models"
	"strconv"
	"time"

	"gorm.io/gorm/clause"
)

func LoadCSV(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	_, err = reader.Read() // skip header
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// Parse fields
		orderID := record[0]
		productID := record[1]
		customerID := record[2]
		productName := record[3]
		category := record[4]
		region := record[5]
		dateOfSale, err := time.Parse("2006-01-02", record[6])
		if err != nil {
			// handle error, skip or log
			continue
		}
		quantitySold, err := strconv.Atoi(record[7])
		if err != nil {
			continue
		}
		unitPrice, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			continue
		}
		discount, err := strconv.ParseFloat(record[9], 64)
		if err != nil {
			continue
		}
		shippingCost, err := strconv.ParseFloat(record[10], 64)
		if err != nil {
			continue
		}
		paymentMethod := record[11]
		customerName := record[12]
		customerEmail := record[13]
		customerAddress := record[14]

		// Insert or update Customer
		customer := models.Customer{
			ID:      customerID,
			Name:    customerName,
			Email:   customerEmail,
			Address: customerAddress,
		}
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&customer)

		// Insert or update Product
		product := models.Product{
			ID:       productID,
			Name:     productName,
			Category: category,
		}
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&product)

		// Insert or update Order
		order := models.Order{
			ID:            orderID,
			CustomerID:    customerID,
			Region:        region,
			DateOfSale:    dateOfSale,
			PaymentMethod: paymentMethod,
		}
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&order)

		// Insert OrderItem
		orderItem := models.OrderItem{
			ID:           orderID + "_" + productID, // unique id for order item
			OrderID:      orderID,
			ProductID:    productID,
			Quantity:     quantitySold,
			UnitPrice:    unitPrice,
			Discount:     discount,
			ShippingCost: shippingCost,
		}
		db.DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&orderItem)
	}

	log.Println("CSV loaded successfully.")
	return nil
}
