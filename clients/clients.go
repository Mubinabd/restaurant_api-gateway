package clients

import (
	"log/slog"

	pb "github.com/Mubinabd/restaurant_api-gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	RestaurantClient  pb.RestaurantServiceClient
	OrderClient       pb.OrderServiceClient
	MenuClient        pb.MenuServiceClient
	PaymentClient     pb.PaymentServiceClient
	ReservationClient pb.ReservationServiceClient
}

func NewClients() *Clients {
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("error:", err)
	}
	restC := pb.NewRestaurantServiceClient(conn)
	orderC := pb.NewOrderServiceClient(conn)
	menuC := pb.NewMenuServiceClient(conn)
	reservationC := pb.NewReservationServiceClient(conn)
	return &Clients{
		RestaurantClient:  restC,
        OrderClient:       orderC,
        MenuClient:        menuC,
        // PaymentClient:     paymentC,
        ReservationClient: reservationC,
	}
}
