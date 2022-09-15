package main

import "fmt"

//Struct
type BankAccount struct {
	name     string
	password string
}

//Login Module
func login() {
	var username string
	var userpassword string
	var Select string
	fmt.Println("+++Welcome ATM Machine Phase 1+++")
	fmt.Scan("Please Select option", Select)
	fmt.Println("I->Login")
	fmt.Println("C->Create Bank Account")
	fmt.Println("q->Quit")
	fmt.Scan(&Select)
	x := BankAccount{
		name:     "Sandhya",
		password: "Saisandhya@0708",
	}
	if Select == "I" {
		fmt.Println("Please enter UserName")
		fmt.Scan(&username)
		fmt.Println("Please enter Password")
		fmt.Scan(&userpassword)
		if username == x.name && userpassword == x.password {
			fmt.Println("Welcome to", x.name)
			fmt.Println("WM->Withdraw Money.")
			fmt.Println("DP->Deposit money..")
			fmt.Println("RB->Request balance.")
			fmt.Println("Quit->Quit the program.")
		} else {
			fmt.Println("UnAthorized User")
			fmt.Println("Please Visit Home Page")
		}
	}

	if Select == "C" {
		fmt.Println("Create Bank Account")
		fmt.Println("Please Create Account and Fill All Details")
	}
	if Select == "q" {
		fmt.Println("Exit the Page")
		fmt.Println("Please Visit Login Page")

	}
}
func BankOperations() {

	var choose string
	fmt.Scan(&choose)
	if choose == "WM" {
		fmt.Println("Please withDraw Money")
	}
	if choose == "DP" {
		fmt.Println("Deposit Money")
	}
	if choose == "RB" {
		fmt.Println("Request for Money")
	}
	if choose == "Quit" {
		fmt.Println("leave the Page")
	}
}
func main() {
	login()
	BankOperations()
}
