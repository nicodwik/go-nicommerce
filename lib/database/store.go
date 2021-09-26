package database

import (
	"mini-project-acp12/config"
	"mini-project-acp12/models"
)

func InsertStore(store *models.Store) (interface{}, error) {
	if err := config.DB.Create(&store).Error; err != nil {
		return nil, err
	}

	return store, nil
}
