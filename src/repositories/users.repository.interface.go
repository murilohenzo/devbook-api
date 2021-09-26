package repositories

import "api_resources/src/models"

type IRepository interface {
	Create(user *models.User) (*models.User, error)
}
