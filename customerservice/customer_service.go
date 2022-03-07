package customerservice

import (
	"github.com/armuh16/patternfacade/hr"
	"github.com/armuh16/patternfacade/sales"
)

type Ticket struct {
	User    string
	Type    string
	Message string
}

type Response struct {
	Status string
}

func HandleTicket(t Ticket) Response {
	switch t.Type {
	case "sales":
		response := sales.ProcessTicket(t.Message)
		return Response{Status: response}
	case "hr":
		response := hr.ProcessTicket(t.Message)
		return Response{Status: response}
	}
	return Response{Status: "customer service is busy"}
}
