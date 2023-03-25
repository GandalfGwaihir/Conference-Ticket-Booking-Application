package main


import (
	"fmt"
	"strconv"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)





func main() {
	

	greetUsers()

	for {

		var firstName, lastName, email, userTickets = getUserInput()
		var isValidName, isValidEmail, isValidTicketNumber = validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			bookTicket( userTickets,  firstName, lastName, email)

			//call the function firstNames
			var firstNames = getFirstNames()
			fmt.Printf("These first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Sorry, we are sold out of tickets")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}


		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and, %v tickets are remaining \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		
		firstNames = append(firstNames, booking["firstName"])

	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket( userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}



// city := "London"

			// switch city {
			//     case "New York":
			// 		// execute code for booking New York conference tickets
			// 	case "Singapore", "Hong Kong":
			// 		// execute code for booking Singapore & Hong Kong conference tickets
			// 	case "London", "Berlin":
			// 		// execute code for booking London & Berlin conference tickets

			// 	case "Mexico City":
			// 		// execute code for booking Mexico City conference tickets
			// 	default:
			// 		fmt.Print("No valid city selected")
			// }