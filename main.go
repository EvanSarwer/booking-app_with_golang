package main

import "fmt"

func main() {
	//var conferenceName = "Go Conference"          // variable declaration with initialize value is ok without specifying the type
	// conferenceName := "Go Conference"            // syntactic sugar in Go
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //unsigned integer can't be negative
	//define array
	//var bookings = [50]string{"Evan", "Sarwer", "Arif"}           // declare an array of 50 elements and assign values
	//var bookings [50]string                                   // declare an array of 50 elements without assigning values
	//define slice
	//var bookings = []string{"Evan", "Sarwer", "Arif"}           // declare a slice of 3 elements and assign values
	var bookings []string								   // declare a slice of 50 elements without assigning values

	fmt.Printf("nonferenceName is %T and conferenceTickets is %T\n", conferenceName, conferenceTickets)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets hete to attend")


	



	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	//fmt.Println(&userName)        // pointer to the memory location of the variable
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	//bookings[0] = firstName + " " + lastName          // assign value to the first element of the array
	bookings = append(bookings, firstName + " " + lastName)     // append value to the slice

	

	fmt.Printf("Thak you %v %v, for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	fmt.Printf("There are all our bookings: %v", bookings)

}
