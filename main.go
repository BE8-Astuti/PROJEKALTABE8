package main

import (
	"be8/gorm/config"
	"be8/gorm/datastore"
	"be8/gorm/entities"
	"fmt"

	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	dbConn = config.InitDB()
	InitialMigration()
}

func InitialMigration() {

	dbConn.AutoMigrate(&entities.Book{})
	dbConn.AutoMigrate(&entities.User{})
	dbConn.AutoMigrate(&entities.Rented_Book{})
}

func main() {
	again := true
	for again == true {
		fmt.Println("MENU:")
		fmt.Println("1 - Register")
		fmt.Println("2 - Login")
		fmt.Println("3 - Lihat daftar buku")

		var pilihan string
		fmt.Println("\nMasukkan pilihan anda:")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":

			user := entities.User{}

			fmt.Println("\nMasukkan Nama:")
			fmt.Scanln(&user.Name)
			fmt.Println("Masukkan Phone:")
			fmt.Scanln(&user.Phone)
			fmt.Println("Masukkan Email:")
			fmt.Scanln(&user.Email)
			fmt.Println("Masukkan Password:")
			fmt.Scanln(&user.Password)
			fmt.Println("Masukkan Address:")
			fmt.Scanln(&user.Address)

			result := datastore.CreateUser(dbConn, user)
			if result != nil {
				fmt.Println("CREATE ACCOUNT FAIL")
			}

		case "2":

		case "3":

		case "5":

		case "6":

		case "7":

		case "8":

		}

		fmt.Println("\n\nNOTE:")
		fmt.Println("[1] Back to MENU")
		fmt.Println("[2] LogOut")
		fmt.Println("\nMasukkan pilihan anda:")
		fmt.Scanln(&pilihan)
		if pilihan == "2" {
			again = false
			fmt.Println("\n\t*-* *-* *-* THANK YOU *-* *-* *-*")
			fmt.Println("\t\t\tCreated by TIM 1")

		}
	}
}
