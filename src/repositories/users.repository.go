package repositories

import (
	"api_resources/src/database"
	"api_resources/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func New() *UserRepository {
	db := database.InitDb()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &UserRepository{Db: db}
}

//Create user
func Create(User *models.User) (err error) {
	err = 	New().Db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll get users
func FindAll(db *gorm.DB, User *[]models.User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// FindById get user by id
func FindById(db *gorm.DB, User *models.User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//Update user
func Update(db *gorm.DB, User *models.User) (err error) {
	db.Save(User)
	return nil
}

//Delete user
func Delete(db *gorm.DB, User *models.User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}