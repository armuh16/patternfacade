package main

import (
	"fmt"

	"github.com/armuh16/patternfacade/customerservice"
)

func main() {
	response := customerservice.HandleTicket(customerservice.Ticket{
		User:    "armuh",
		Type:    "marketing sales",
		Message: "message for marketing officer",
	})
	fmt.Println(response.Status)
	response1 := customerservice.HandleTicket(customerservice.Ticket{
		User:    "arif muhammad",
		Type:    "hr",
		Message: "message for hr",
	})
	fmt.Println(response1.Status)
}
