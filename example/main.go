package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ivangurin/tinkoff-invest-client-go/pkg/client"
	"github.com/ivangurin/tinkoff-invest-client-go/pkg/investapi"
)

func main() {

	// Получение клиента
	client, err := client.NewClient(os.Getenv("TINKOFF_INVEST_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Отпрвка запроса на получение счетов
	response, err := client.UsersService.GetAccounts(context.Background(), &investapi.GetAccountsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// Вывод результата на экран
	for _, account := range response.Accounts {
		fmt.Printf("%+v\n", account)
	}

}
