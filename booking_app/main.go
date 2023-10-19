package main

import "fmt"

func main() {

	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var reminingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, reminingTickets)
	fmt.Println("Get your ticket here to attend.")
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remindingTickers is %T \n", conferenceName, conferenceTickets, reminingTickets)

	var userName string
	var userTickets int

	userName = "Tyga"
	userTickets = 40
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

	//// Println
	// fmt.Println("Welcome to", conferenceName, "Booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and", reminingTickets)
	// fmt.Println("Get your ticket here to attend.")

	// fmt.Println("Here we go with variables")

}
