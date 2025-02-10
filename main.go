package main

import "fmt"

func main() {
	//var conferenceName = "Go Conference"          // variable declaration with initialize value is ok without specifying the type
	// conferenceName := "Go Conference"            // syntactic sugar in Go 
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50                  //unsigned integer can't be negative

	fmt.Printf("nonferenceName is %T and conferenceTickets is %T\n", conferenceName, conferenceTickets)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets hete to attend")

	var userName string
	var userTickets int

}
