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

func LihatUsers(db *gorm.DB) []entities.User {
	var users []entities.User
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}
