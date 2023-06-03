package client

import (
	"Projects/store/store_api_gateway/config"
	"Projects/store/store_api_gateway/genproto/order_service"
	"Projects/store/store_api_gateway/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	//userservice
	UserService() user_service.UserServiceClient
	//orderservice
	OrderService() order_service.OrderServiceClient
}

type grpcClients struct {
	userService  user_service.UserServiceClient
	orderService order_service.OrderServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	//connecting to user service
	connUserService, err := grpc.Dial(
		cfg.UserServiceHost+cfg.UserServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	// connecting to order service
	connOrderService, err := grpc.Dial(
		cfg.OrderServiceHost+cfg.OrderServicePort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:  user_service.NewUserServiceClient(connUserService),
		orderService: order_service.NewOrderServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) OrderService() order_service.OrderServiceClient {
	return g.orderService
}
