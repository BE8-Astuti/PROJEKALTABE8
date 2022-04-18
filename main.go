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
		fmt.Println("\n==========================================")
		fmt.Println(">< >< >< ALTA RENT BOOK APPS>< >< ><\n\t\t^ARB^")
		fmt.Println("==========================================")
		fmt.Println("1 - Sign Up")
		fmt.Println("2 - Login")
		fmt.Println("3 - All Available Book")

		var pilihan string
		fmt.Println("\nInput your Choice, please: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":

			user := entities.User{}

			fmt.Println("\n==========================================")
			fmt.Println(">< >< >< REGISTER NEW ACCOUNT >< >< ><")
			fmt.Println("==========================================")

			fmt.Println("\nName:")
			fmt.Scanln(&user.Name)
			fmt.Println("Phone:")
			fmt.Scanln(&user.Phone)
			fmt.Println("Email:")
			fmt.Scanln(&user.Email)
			fmt.Println("Password:")
			fmt.Scanln(&user.Password)
			fmt.Println("Address:")
			fmt.Scanln(&user.Address)

			result := datastore.CreateUser(dbConn, user)
			if result != nil {
				fmt.Println("\n\n=========================================")
				fmt.Println("!?!?!?!!?? CREATE ACCOUNT FAIL ??!!!?!?!?!")
				fmt.Println("=========================================")
			} else {
				fmt.Println("\n\n=========================================")
				fmt.Println("^^^^^^^^ CREATE ACCOUNT SUCCESS ^^^^^^^^")
				fmt.Println("=========================================")
			}

		case "2":

			user := entities.User{}
			fmt.Println("\nMasukkan Email : ")
			fmt.Scanln(&user.Email)
			fmt.Println("\nMasukkan Password : ")
			fmt.Scanln(&user.Password)
			result := datastore.GetUserbyPassword(dbConn, user.Password)

			for _, v := range result {
				confirm := ""
				for confirm != "0" {
					fmt.Println("\n\n\n=========================================================")
					fmt.Println("\t  WELCOME to ALTA RENT BOOK (ARB) ")
					fmt.Printf("\t\t\t %s ^-^", v.Name)
					fmt.Println("\n=========================================================")

					fmt.Println("\n\nWHAT DO YOU NEED ?")
					fmt.Println("\n[A] Profile")
					fmt.Println("[B] My Book")
					fmt.Println("[C] Available Book")
					fmt.Println("[D] Borrow Book")
					fmt.Println("[E] Rented Book")
					fmt.Println("[F] Update Book")
					fmt.Println("[G] Add NEW Book")
					fmt.Println("[H] Return Book")
					fmt.Println("[0] LOG OUT")
					fmt.Printf("\nConfirm %s, please: ", v.Name)
					confirm := ""
					fmt.Scanln(&confirm)
					// Back := true

					switch confirm {
					case "A":

						for confirm != "d" {

							fmt.Println("\n\n============================================")
							fmt.Println("\t >< >< >< MENU PROFILE >< >< ><")
							fmt.Println("============================================")
							fmt.Println("[a] Read Profile")
							fmt.Println("[b] Update Profile")
							fmt.Println("[c] Non Active ID")
							fmt.Println("[d] MENU LOG IN")
							fmt.Println("=============================")
							fmt.Printf("\nWhat do you like %s? Confirm,please: ", v.Name)
							fmt.Scanln(&confirm)
							if confirm == "a" {
								fmt.Println("======================================")
								fmt.Printf("\t<> <> PROFIL of %s <> <>", v.Name)
								fmt.Println("\n======================================")
								fmt.Println("\nUser ID\t: ", v.ID)
								fmt.Println("Email\t: ", v.Email)
								fmt.Println("Phone\t: ", v.Phone)
								fmt.Println("Address\t: ", v.Address)
								fmt.Println("=============================")

							} else if confirm == "b" {

								var email string
								var address string

								var result error
								fmt.Println("=============================")
								fmt.Println("\n\nMenu Update Profile:")
								fmt.Println("\n[1]  EMAIL")
								fmt.Println("[2]  ADDRESS")
								fmt.Println("=============================")

								fmt.Printf("\nConfirm %s, please: ", v.Name)
								confirm := ""
								fmt.Scanln(&confirm)
								// for confirm != 2 {
								if confirm == "1" {
									fmt.Println("\nInput NEW EMAIL : ")
									fmt.Scanln(&email)
									result = datastore.UpdateAccountEmail(dbConn, v.Password, email)

								} else {

									fmt.Println("Input NEW ADDRESS : ")
									fmt.Scanln(&address)
									result = datastore.UpdateAccountAddress(dbConn, v.Password, address)
								}

								if result == nil {
									fmt.Println("\n\n======================================================")
									fmt.Println(">> >> >> UPDATE SUCCESS ~~~~~ UPDATE SUCCESS << << <<")
									fmt.Println("======================================================")
								} else {
									fmt.Println("\n===================================================")
									fmt.Println(">> >> >> UPDATE FAIL ~~~~~ UPDATE FAIL << << <<")
									fmt.Println("=====================================================")
								}

							} else if confirm == "c" {

								account := entities.User{}
								fmt.Println("=========================================================")
								fmt.Println("WARNING!!! (Are You Sure to Non-Active Your Account?)\n [1] YES \n [2] NO")
								fmt.Println("=========================================================")
								fmt.Printf("\nConfirm %s, please: ", v.Name)
								choice := ""
								fmt.Scanln(&choice)
								if choice == "1" {
									fmt.Println("Enter your USER ID : ")
									fmt.Scanln(&account.ID)
								} else {
									main()
								}

								result := datastore.DeleteUser(dbConn, account.ID)
								if result == nil {
									fmt.Println("=================================================")
									fmt.Println("\n\t~~~~~~~~ ACCOUNT NOT ACTIVE ANYMORE ~~~~~~~~")
									fmt.Println("=================================================")
								} else {
									fmt.Println("=================================================")
									fmt.Println("\n\t~~~~~~~~~~ NOT ANY ACCOUNT ~~~~~~~~~~~~")
									fmt.Println("=================================================")
								}
							}

						}

					case "B":

						result := datastore.GetRentedBookbyUserid(dbConn, v.ID)
						fmt.Println("================================================================================================================")
						fmt.Printf("\t\t\t\t<><><> LIST of %s's BOOKS <><><>", v.Name)
						fmt.Println("\n================================================================================================================")
						for _, v := range result {
							fmt.Printf("\n[ %v ]\t%s\t\t%s\t\t%s\t%s", v.Book_id, v.Judul, v.Pengarang, v.CreatedAt, v.UpdatedAt)
						}
						fmt.Println("\n================================================================================================================")
						if len(result) < 1 {
							fmt.Println("=================================================")
							fmt.Println("\t\t~~~ SORRY, NOT ANY BOOK ~~~")
							fmt.Println("=================================================")
						}

					case "C":

						result := datastore.GetBook(dbConn)
						fmt.Println("\n=========================================================================================================================")
						fmt.Println("\t\t\t\t\t><><><>< List of AVAILABLE BOOK ><><><><><\n\t\t\t\t\t\t  ALTA RENT BOOK (^ARB^")
						fmt.Println("========================================================================================================================")
						for _, v := range result {
							fmt.Printf("\n[ %v ]\t%s\t\t%s\t\t%v\t\t%s\t\t%s", v.ID, v.Judul, v.Pengarang, v.Tahun_terbit, v.CreatedAt, v.UpdatedAt)
						}
						if len(result) < 1 {
							fmt.Println("=================================================")
							fmt.Println("\t ~~~ SORRY, NOT ANY BOOK ~~~")
							fmt.Println("=================================================")
						}
						fmt.Println("\n=========================================================================================================================")

					case "D":
						borrow := entities.Rented_Book{}
						// lender := entities.User{}
						// book := entities.Book{}
						fmt.Println("\nInput ID BOOK:")
						// fmt.Scanln(&book.ID)b
						fmt.Scanln(&borrow.Book_id)
						fmt.Println("\nInput TITLE:")
						fmt.Scanln(&borrow.Judul)

						resultt := datastore.GetRentedBookbyBookid(dbConn, borrow.Book_id)

						if len(resultt) != 0 {
							for _, v := range resultt {
								fmt.Println("\n\n=================================================")
								fmt.Println("  !?!?! SORRY, THE BOOK IS NOT AVAILABLE !?!?!?")
								lenders := datastore.GetUserbyID(dbConn, v.User_id)
								fmt.Printf("\n~ ^ %s ^ has been rented by %s (ID: %v) ~", borrow.Judul, lenders[0].Name, v.ID)
								fmt.Println("\n=================================================")

								fmt.Println("\n\n\nNB:")
								fmt.Printf("[1] Would you like to lending from { %s } ?", lenders[0].Name)
								fmt.Println("\n[2] Find other books ^-^")

								fmt.Println("\nConfirm, please: ")
								confirm := ""
								fmt.Scanln(&confirm)
								if confirm == "1" {

									fmt.Println("\nInput My USER ID:")
									fmt.Scanln(&borrow.User_id)
									lending := datastore.UpdateRentedBook(dbConn, borrow.Book_id, borrow.User_id)
									if lending == nil {
										fmt.Println("\n================================================================")
										fmt.Printf("\t\t\tCONGRATULATION ^-^\n\tYou have just rented %s (ID Book: %v) from { %s }\n\n>->>>->>>-> RENT BOOK SUCCESS ~~~ RENT BOOK SUCCESS <-<<<-<<<-<", borrow.Judul, borrow.Book_id, lenders[0].Name)
										fmt.Println("\n================================================================")
									} else {
										fmt.Println("\n\n========================================")
										fmt.Println(" ~~~~~ SORRY, RENT BOOK FAIL ~~~~~~")
										fmt.Println("=======================================")
									}
								}
							}

						} else {
							users := datastore.GetUserbyPassword(dbConn, v.Password)
							books := datastore.GetBookbyid(dbConn, borrow.Book_id)
							borrow.User_id = users[0].ID
							borrow.Judul = books[0].Judul
							borrow.Pengarang = books[0].Pengarang
							borrow.Tahun_terbit = books[0].Tahun_terbit
							borrows := datastore.CreateBorrow(dbConn, borrow)

							updateavailablebook := datastore.DeleteAvBook(dbConn, borrow.Book_id)
							if borrows == nil && updateavailablebook == nil {
								fmt.Println("\n========================================================================")
								fmt.Printf("\t\t\tCONGRATULATION ^-^\n\t\t%s has rented %s (IDBook: %v)\n\n>->>>->>>-> RENT BOOK SUCCESS ~~~ RENT BOOK SUCCESS <-<<<-<<<-<", v.Name, borrow.Judul, borrow.Book_id)
								fmt.Println("\n========================================================================")
							} else {
								fmt.Println("\n========================================")
								fmt.Println(" ~~~~~ SORRY, RENT BOOK FAIL ~~~~~~")
								fmt.Println("=======================================")
							}
						}
					case "E":

						see := datastore.GetRentedBook(dbConn)
						fmt.Println("\n=========================================================================================================================================")
						fmt.Println("\t\t\t\t\t><><><>< List of RENTED BOOKS ><><><><><\n\t\t\t\t\t\t  ALTA RENT BOOK (^ARB^)")
						fmt.Println("========================================================================================================================================")
						for _, v := range see {
							fmt.Printf("\n[ %v ]\t[ %v ]\t%s\t\t%s\t\t%v\t\t%s\t\t%s", v.ID, v.User_id, v.Judul, v.Pengarang, v.Tahun_terbit, v.CreatedAt, v.UpdatedAt)
						}
						if len(result) < 1 {
							fmt.Println("=================================================")
							fmt.Println("\t ~~~ SORRY, NOT ANY BOOK ~~~")
							fmt.Println("=================================================")
						}
						fmt.Println("\n========================================================================================================================================")

					case "F":

						var pengarang string
						var judul string

						var result error
						bookss := entities.Book{}
						fmt.Println("\nInput ID BOOK : ")
						fmt.Scanln(&bookss.ID)
						fmt.Println("=============================")
						fmt.Println("CHOOSE: ")
						fmt.Println("\n[1] UPDATE Title")
						fmt.Println("[2] UPDATE Authors")
						fmt.Println("=============================")

						fmt.Printf("\nConfirm %s, please: ", v.Name)
						confirm := ""
						fmt.Scanln(&confirm)
						if confirm == "1" {
							fmt.Println("\nInput NEW TITLE : ")
							fmt.Scanln(&judul)
							result = datastore.UpdateBookJudul(dbConn, bookss.ID, judul)

						} else {

							fmt.Println("Input NEW AUTHORS : ")
							fmt.Scanln(&pengarang)
							result = datastore.UpdateBookPengarang(dbConn, bookss.ID, pengarang)

						}

						if result == nil {
							fmt.Println("\n==================================================")
							fmt.Println("\t>> >> >> UPDATE SUCCESS << << <<")
							fmt.Println("==================================================")
						} else {
							fmt.Println("=================================================")
							fmt.Println("\t>> >> >> UPDATE FAIL << << <<")
							fmt.Println("=================================================")
						}

					case "G":
						books := entities.Book{}

						fmt.Println("\nInput Title: ")
						fmt.Scanln(&books.Judul)
						fmt.Println("Input Authors: ")
						fmt.Scanln(&books.Pengarang)
						fmt.Println("Input Year: ")
						fmt.Scanln(&books.Tahun_terbit)

						result := datastore.AddBook(dbConn, books)
						if result != nil {
							fmt.Println("=========================================")
							fmt.Println("  !?!?!?!!?? ADD BOOK FAIL ??!!!?!?!?!")
							fmt.Println("=========================================")
						} else {
							fmt.Println("=========================================")
							fmt.Println(" ><><><><> ADD BOOK  SUCCESS <><><><><><")
							fmt.Println("=========================================")
						}

					case "H":
						bbok := entities.Book{}
						fmt.Println("\n=============================")
						fmt.Printf("Have a nice day %s !!!", v.Name)
						fmt.Println("\n[a] Return BOOK to ARB APPS")
						fmt.Println("[b] Return BOOK to LENDER")
						fmt.Println("=============================")
						fmt.Printf("\nConfirm %s, please: ", v.Name)
						fmt.Scanln(&confirm)

						if confirm == "a" {
							fmt.Println("\nInput Title: ")
							fmt.Scanln(&bbok.Judul)
							fmt.Println("Input Authors: ")
							fmt.Scanln(&bbok.Pengarang)
							fmt.Println("Input Year: ")
							fmt.Scanln(&bbok.Tahun_terbit)

							result := datastore.AddBook(dbConn, bbok)

							if result == nil {
								fmt.Println("\n========================================================================")
								fmt.Printf("\t\t\t\t^-^ ^-^ ^-^\n\t\tYou have just returned %s (Year: %v)\n\n>->>>->>>-> RETURN BOOK SUCCESS ~~~ RENTURN BOOK SUCCESS <-<<<-<<<-<", bbok.Judul, bbok.Tahun_terbit)
								fmt.Println("\n=========================================================================")
							} else {
								fmt.Println("=================================================")
								fmt.Println("\n!?!?!?!!?? RENT BOOK FAIL !?!?!?!!??")
								fmt.Println("=================================================")
							}
						} else {
							borroww := entities.Rented_Book{}
							fmt.Println("\nInput Lender`s USER ID:")
							fmt.Scanln(&borroww.User_id)
							fmt.Println("\nInput ID BOOK:")
							fmt.Scanln(&borroww.Book_id)
							fmt.Println("\nInput TITLE:")
							fmt.Scanln(&borroww.Judul)

							lending := datastore.UpdateRentedBook(dbConn, borroww.Book_id, borroww.User_id)
							if lending == nil {
								lastlenders := datastore.GetUserbyID(dbConn, borroww.User_id)

								fmt.Println("\n=====================================================================")
								fmt.Printf("\t\t^-^ ^-^ ^-^\n\tYou have just returned %s (ID Book: %v) to %s \n\n>->>>->>>-> RENTURN BOOK SUCCESS ~~~ RENTURN BOOK SUCCESS <-<<<-<<<-<", borroww.Judul, borroww.Book_id, lastlenders[0].Name)
								fmt.Println("\n=====================================================================")
							} else {
								fmt.Println("\n\n========================================")
								fmt.Println(" ~~~~~ SORRY, NOT FOUND ID BOOK ~~~~~~")
								fmt.Println("=======================================")
							}
						}

					}

					if confirm == "0" {
						main()

					}

				}

			}
			if result != nil {
				fmt.Println("\n\n=================================================")
				fmt.Println("  ~~~~~~~~SORRY, ACCOUNT NOT ACTIVE  ~~~~~~~~")
				fmt.Println("=================================================")
			}

		case "3":
			result := datastore.GetBook(dbConn)
			fmt.Println("\n=========================================================================================")
			fmt.Println("\t\t\t><><><>< List of AVAILABLE BOOK ><><><><><\n\t\t\t\t  ALTA RENT BOOK APPS")
			fmt.Println("=========================================================================================")
			for _, v := range result {
				fmt.Printf("\n%v\t%s\t\t%s\t\t%v\t\t%s", v.ID, v.Judul, v.Pengarang, v.Tahun_terbit, v.CreatedAt)

			}
			if len(result) < 1 {
				fmt.Println("=================================================")
				fmt.Println("\t\t~~~ SORRY, NOT ANY BOOK ~~~")
				fmt.Println("=================================================")
			}
			fmt.Println("\n=========================================================================================")
			// result := datastore.GetRentedBook(dbConn)
			// fmt.Println("=================================================")
			// fmt.Printf("\n\t\t<><> LIST RENTED BOOK <><>")
			// fmt.Println("\n================================================")
			// for _, v := range result {
			// 	fmt.Printf("\n\t%s\t\t%s\t\t%v\t\t%s", v.Judul, v.Pengarang, v.Tahun_terbit, v.CreatedAt)
			// }
			// if len(result) < 1 {
			// 	fmt.Println("=================================================")
			// 	fmt.Println("\t~~~ SORRY, NOT ANY BOOK ~~~")
			// 	fmt.Println("\n================================================")
			// }

		}
		fmt.Println("\n\nNOTE:")
		fmt.Println("[1] Back to ARB APPS")
		fmt.Println("[2] Close ARB APPS")
		fmt.Println("\nConfirm, please : ")
		fmt.Scanln(&pilihan)
		if pilihan == "2" {
			again = false
			fmt.Println("\n=================================================================================")
			fmt.Println("\t>->>>->>>-> THANK YOU ~~~~~ THANK YOU  ~~~~~ THANK YOU <-<<<-<<<-<")
			fmt.Println("\t\t\t  Alta Immersive Progam BE 8")
			fmt.Println("=================================================================================")

		}

	}

}
