package datastore

import (
	"be8/gorm/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser entities.User) error {
	result := db.Create(&newUser)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

func GetUserbyPassword(db *gorm.DB, password string) []entities.User {
	var users []entities.User
	tx := db.Find(&users, "password = ?", password)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}

func GetUserbyID(db *gorm.DB, ID uint) []entities.User {
	var users []entities.User
	tx := db.Find(&users, "id = ?", ID)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}

func UpdateMyBook(db *gorm.DB, id uint, judul string) error {
	result := db.Model(&entities.User{}).Where("id = ?", id).Update("judul", judul)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
func UpdateAccountEmail(db *gorm.DB, password, email string) error {
	result := db.Model(&entities.User{}).Where("password = ?", password).Update("email", email)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}

func UpdateAccountAddress(db *gorm.DB, password, address string) error {
	result := db.Model(&entities.User{}).Where("password = ?", password).Update("address", address)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}

func DeleteUser(db *gorm.DB, id uint) error {
	var user []entities.User
	result := db.Delete(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found, update failed")
	}
	return nil
}
