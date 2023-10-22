package main

import (
	"booking_app/helper"
	"fmt"
	"strings"
	// import the helper from my booking_app module
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greatUsers(remainingTickets)

	fmt.Println("Welcome to", conferenceName, "Booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are available")
	fmt.Println("Get your ticket here to attend.")

	for {
		// get userInput
		firstName, lastName, email, userTickets := getUserInput()

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidationUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// here we call function to book ticket
			bookTicket(firstName, lastName, email, userTickets)
			// here we call function print first name
			firstNames := firstNames()
			fmt.Printf("The first name of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The conference is booked out, come next year")
				break

			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

}

func greatUsers(remainingTickets uint) { // function with params
	fmt.Printf("Welcome to  %v Booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
}

func firstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) //This will splite it
		var firstName = names[0]
		firstNames = append(firstNames, firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {

	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you for booking %v you will receive email confirmation on %v\n", userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
