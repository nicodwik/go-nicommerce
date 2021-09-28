package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/models"
)

func InsertTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := config.DB.Create(&transaction).Error; err != nil {
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
