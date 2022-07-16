package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const confernceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUser(conferenceName, confernceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their names
		fmt.Println("Enter first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userTickets)
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidTicketNumber && isValidEmail {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence id booked out.come next year.")
				break

			}

		} else {

			if !isValidName {

				fmt.Printf("First Name or the last name is too short \n")

			}

			if !isValidEmail {

				fmt.Printf("Invalid email \n")

			}

			if !isValidTicketNumber {

				fmt.Printf("Invalid ticket number \n")

			}
		}

	}
}
func greetUser(confName string, confTickets int, remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available. \n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
func getFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func validateUserInput(firstName string, lastName, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber

}
