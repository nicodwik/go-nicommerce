package database

import (
	"go-nicommerce/config"
	"go-nicommerce/models"
)

func InsertTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := config.DB.Create(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetTransactionByOrderCode(orderCode string) (*models.Transaction, error) {
	var transaction models.Transaction

	if err := config.DB.Where("order_code = ?", orderCode).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func UpdateTransactionData(transaction *models.Transaction) (*models.Transaction, error) {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func InsertTransactionProducts(transactionProducts *[]models.TransactionProduct) ([]models.TransactionProduct, error) {
	if err := config.DB.Create(&transactionProducts).Error; err != nil {
		return nil, err
	}

	return *transactionProducts, nil
}
