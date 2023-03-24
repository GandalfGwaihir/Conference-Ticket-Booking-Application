package main

import "fmt"


func main(){
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T \n",conferenceTickets, remainingTickets,conferenceName)


	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and, %v tickets are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")


	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//Ask the user for  their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) 
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName) 
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email) 
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets) 

	
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("slice type: %T \n", bookings)
	fmt.Printf("slice length: %v \n", len(bookings))
	
	remainingTickets = remainingTickets - userTickets
	
	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v \n", bookings)

}