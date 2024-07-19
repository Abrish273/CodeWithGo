package main
// 02:19:00
import (
	"fmt"
	"strings"
)
//  package level variables

const conferenceTickets int = 50
var conferenceName string = "Go conference"
var remainingTickets int = 10


func main() { 
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	greetUser(conferenceName, conferenceTickets, remainingTickets)
	
	
	// var userTickets float32
	
	// Types of loops
	// Infinite loops --- for {}
	// for {
		// Pointers are special types of variables that store the memory address of another variable. 
		// %v - is value and %T - is type
    var bookings [] string
	firstName, lastName, email, Tickets := getUserInput()

	for remainingTickets > 0 && len(bookings) < 50 { //iterate until this condition is true
		
		fmt.Println("Enter your first name: ")
		 isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, email, Tickets, remainingTickets)

		// isValidCity := city == "Nairobi" && city == "Mombasa" 

		// bookings[0] = firstName + " " + lastName
		if isValidName && isValidEmail && isValidTicket {
			bookTicket(remainingTickets, Tickets, firstName, lastName, email)
			
			// Slices - slice is an abstraction of an array
			// Slices are more flexible and powerful : variable length or get a sub-array of its own
			// Slices are also index based and have a size , but is resized when needed.
			
			bookings = append(bookings, firstName + " " + lastName)
			fmt.Printf("These are the names of the bookings: %v \n", bookings)
			
			firstNames := FirstNames(bookings)
		    fmt.Printf("These are the first names of the bookings: %v \n", firstNames)


			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
	} else {
		if !isValidName {
			fmt.Println("Please enter a valid name")
		}else if !isValidEmail {
			fmt.Println("Please enter a valid email")
		} else if !isValidTicket {
			fmt.Println("Please enter a valid name")
        } else {
			fmt.Println("Something went wrong. Please try again")
		}
		// fmt.Printf("the available number of tikets are only %v and you can't book %v number of tickets\n", remainingTickets, Tickets)
		// break
		//continue // makes the loop to continue or to re run again
	}
	
	// var city string

	// switch city {
	// 	case "Nairobi":
	// 		fmt.Println("Welcome to Nairobi")
	// 	case "Mombasa", "Addis Ababa":
	// 		fmt.Println("Welcome to Mombasa or Addis Ababa")
	// 	case "Kisumu":
	// 		fmt.Println("Welcome to Kisumu")
	// 	default:
	// 		fmt.Println("Welcome to the conference")
	// }

   // == , !=, <, <=, >, >=
	}
}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to the %v \n", conferenceName)
	fmt.Printf("We have %v tickets available for the conference and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here . . .")
}

func FirstNames(bookings []string) []string {
        firstNames := [] string {}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		return firstNames
}

func validateUserInput(firstName string, lastName string, email string, Tickets int, remainingTickets int) (bool, bool, bool){
	var isValidName bool = len(firstName) > 2 && len(lastName) > 2
	// isValidName := len(firstName) > 2 && len(lastName) > 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicket bool = Tickets > 0 && Tickets <= remainingTickets
	
	return isValidName, isValidEmail, isValidTicket
}


func getUserInput () (string, string, string, int){
	    var firstName string
		var lastName string
		var email string
		var Tickets int
		// var city string
		// var bookings [50]string // [50] is an array

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your Email: ")
		fmt.Scan(&email)
		fmt.Println("Enter your Tickets: ")
		fmt.Scan(&Tickets)

		return firstName, lastName, email, Tickets
}


func bookTicket(remainingTickets int, Tickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - Tickets
	// fmt.Println("Get your tickets here.", firstName, lastName ," with email of ", email, " and ", Tickets, " tickets")
	fmt.Printf("Thank you %v %v for buying %v tickets and we will send a confirmation message to %v email address\n", firstName, lastName, Tickets, email)
	fmt.Println("remaining Tickets are \n", remainingTickets )
}