# Клиент для Тинькофф инвестиций
Библиотека предназаначена для работы с api Тинькофф инвестиций по grpc протоколу для языка Go. Файлы клиента сгненеры на основе proto-файлов https://github.com/Tinkoff/investAPI/tree/main/src/docs/contracts.

## Установка
```bash
go get -u github.com/ivangurin/tinkoff-invest-client-go
```

## Получение токена
Получить токен можно тут https://www.tinkoff.ru/invest/settings/api.

## Пример использования
```go
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
```

## Лицензия
[MIT](https://choosealicense.com/licenses/mit/)