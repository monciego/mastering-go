package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickers = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickers, "are still available.")
	fmt.Println("Get your tickets here to attend")
}