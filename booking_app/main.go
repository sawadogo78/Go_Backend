package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Println("Welcome to", conferenceName, "Booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets)
	fmt.Println("Get your ticket here to attend.")

	for {
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

		if userTickets < remainingTickets {

			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("The whole array : %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice Type: %T\n", bookings)
			fmt.Printf("Slice length:%v\n", len(bookings))

			fmt.Printf("User %v %v booked % v tickets\n", firstName, lastName, userTickets)
			fmt.Printf("Thank you for booking %v you will receive email confirmation on %v\n", userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				var firstName = names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("All firstNames are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("The conference is booked out, come next year")
				break

			}
		} else {
			fmt.Printf("We only have %v remaining tickets, so can't book %v tickets\n", remainingTickets, userTickets)
		}

	}
}
