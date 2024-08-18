package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var userTickets uint
	var remainingTickets uint = 50
	var firstName string
	var lastName string
	var email string

	fmt.Printf("Welcome to %v Ticketing System\n", conferenceName)
	fmt.Print("Get your conference tickets here\n")
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

