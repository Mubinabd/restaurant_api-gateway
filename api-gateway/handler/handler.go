package handler

import (
	"github.com/Mubinabd/restaurant_api-gateway/clients"

	// pb "github.com/Mubinabd/restaurant_api-gateway/genproto"
	// "google.golang.org/grpc"
)

type HandlerStruct struct {
	Clients  clients.Clients
}

func NewHandlerStruct() *HandlerStruct {
	return &HandlerStruct{
		Clients:*clients.NewClients(),   
	}
}
