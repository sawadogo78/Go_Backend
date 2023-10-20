package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var reminingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "Booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", reminingTickets)
	fmt.Println("Get your ticket here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Enter your first name :")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name :")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email address :")
	fmt.Scan(&email)
	fmt.Printf("Enter your number of tickets :")
	fmt.Scan(&userTickets)

	reminingTickets = reminingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array : %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice Type: %T\n", bookings)
	fmt.Printf("Slice length:%v\n", len(bookings))

	fmt.Printf("User %v %v booked % v tickets\n", firstName, lastName, userTickets)
	fmt.Printf("Thank you for booking %v you will receive email confirmation on %v\n", userTickets, email)
	fmt.Printf("%v tickets remining for %v\n", reminingTickets, conferenceName)

	fmt.Printf("All our bookins %v\n", bookings)
}
