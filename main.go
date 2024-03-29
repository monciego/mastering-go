package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50 

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets )
	fmt.Println("Get your tickets here to attend")

	// array
	// var bookings [50]string
	// slice
	bookings := []string{}

	
	// loops
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)

		/* fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array size: %v\n", len(bookings))
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice size: %v\n", len(bookings))
	*/
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName, lastName, userTickets, email)
		fmt.Printf("There are %v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are : %v\n", firstNames)
	}
}