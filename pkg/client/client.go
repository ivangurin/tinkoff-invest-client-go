package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/ivangurin/tinkoff-invest-client-go/pkg/investapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const prodHost = "invest-public-api.tinkoff.ru"
const sandboxHost = "sandbox-invest-public-api.tinkoff.ru"
const port = "443"

type auth struct {
	Token string
}

func (a auth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + a.Token,
	}, nil
}

func (a auth) RequireTransportSecurity() bool {
	return true
}

type Client struct {
	UsersService       investapi.UsersServiceClient
	InstrumentsService investapi.InstrumentsServiceClient
	MarketDataService  investapi.MarketDataServiceClient
	OperationsService  investapi.OperationsServiceClient
	OrdersService      investapi.OrdersServiceClient
	StopOrdersService  investapi.StopOrdersServiceClient
}

func NewClient(token string) (*Client, error) {
	return new(prodHost, token)
}

func NewSandboxClient(token string) (*Client, error) {
	return new(sandboxHost, token)
}

func new(target string, token string) (*Client, error) {

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: target,
		})),
		grpc.WithPerRPCCredentials(auth{
			Token: token,
		}),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", target, port), opts...)
	if err != nil {
		return nil, err
	}

	client := &Client{
		UsersService:       investapi.NewUsersServiceClient(conn),
		InstrumentsService: investapi.NewInstrumentsServiceClient(conn),
		MarketDataService:  investapi.NewMarketDataServiceClient(conn),
		OperationsService:  investapi.NewOperationsServiceClient(conn),
		OrdersService:      investapi.NewOrdersServiceClient(conn),
		StopOrdersService:  investapi.NewStopOrdersServiceClient(conn),
	}

	return client, nil

}
