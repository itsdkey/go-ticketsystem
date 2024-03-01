package main

import "fmt"

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = conferenceTickets
var bookings []string

func main() {

	greetUsers(conferenceName, remainingTickets)

	for remainingTickets > 0 {
		userName, email, userTickets := getUserInput()
		if !validateUserTickets(userTickets, remainingTickets) {
			continue
		}

		remainingTickets = createBooking(remainingTickets, userTickets, userName, email)

		if remainingTickets == 0 {
			fmt.Printf("No more tickets\n")
			break
		}
	}

}

func greetUsers(conferenceName string, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets still available\n", remainingTickets)
}

func getUserInput() (string, string, uint) {
	var userName string
	var userTickets uint
	var email string

	fmt.Printf("Please provide <name> <email> <number of tickets>:\n")
	fmt.Scanf("%v %v %d", &userName, &email, &userTickets)
	return userName, email, userTickets
}

func validateUserTickets(userTickets, remainingTickets uint) bool {
	isValid := true
	switch {
	case userTickets <= 0:
		fmt.Println("Please provide a postive number of tickets!")
		isValid = false
	case userTickets > remainingTickets:
		fmt.Println("You cannot book that many tickets!")
		isValid = false
	}
	return isValid
}

func createBooking(remainingTickets uint, userTickets uint, userName string, email string) uint {
	remainingTickets -= userTickets
	bookings = append(bookings, userName)

	fmt.Printf("user: %v (%v) - tickets: %v\n", userName, email, userTickets)
	fmt.Printf("bookings: %v\n", bookings)
	fmt.Printf("remaining tickets: %v\n", remainingTickets)
	return remainingTickets
}
