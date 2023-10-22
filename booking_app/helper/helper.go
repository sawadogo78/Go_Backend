package helper // to tell that this helper belongs to main file (should be incorpored there)
import "strings"

func ValidationUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// user input validation logic
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

// export it to make it available for all package in the app, for that we just need to start the function definition by capital letter
