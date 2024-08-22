package main

import ("fmt" 
		"strings")

func main() {
	var conferenceName = "Go Conference"
	var remainingTickets uint = 50
	var userTickets uint
	var firstName string
	var lastName string
	var email string
	var bookings []string
	fmt.Printf("Welcome to %v Ticketing System\n", conferenceName)
	fmt.Print("Get your conference tickets here\n")
	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		enterTickets:
			fmt.Println("Enter number of tickets:")
			fmt.Scan(&userTickets)
		

		if userTickets <= 0 {
			fmt.Println("Please enter a value more than 0.")
			goto enterTickets
		}

		if userTickets > remainingTickets{
			fmt.Printf("We only have %v tickets remaining. You cannot book more than %v tickets\n.", remainingTickets, remainingTickets)
			goto enterTickets
		}
		

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
		firstNames := []string{}
		for _, booking := range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("The conference is booked out. Please visit next year. :)")
			break
		}


	}
}

