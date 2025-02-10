package main

import "fmt"

func main() {
	var conferenceName =  "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	
	fmt.Println("Welcome to",conferenceName, "booking application")
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets hete to attend")
}
