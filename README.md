# Go programming language - Kafka

## Go Packages

- sarama [https://pkg.go.dev/github.com/shopify/sarama](https://pkg.go.dev/github.com/shopify/sarama)
- viper [https://pkg.go.dev/github.com/spf13/viper](https://pkg.go.dev/github.com/spf13/viper)
- fiber [https://pkg.go.dev/github.com/gofiber/fiber/v2](https://pkg.go.dev/github.com/gofiber/fiber/v2)
- gorm [https://pkg.go.dev/gorm.io/gorm](https://pkg.go.dev/gorm.io/gorm)
- gorm MySQL driver [https://pkg.go.dev/gorm.io/driver/mysql](https://pkg.go.dev/gorm.io/driver/mysql)

``` bash
# Install sarama package
go get github.com/Shopify/sarama

# Install viper package
go get github.com/spf13/viper

# Install fiber package
go get github.com/gofiber/fiber/v2

# Install gorm package
go get gorm.io/gorm

# Install gorm MySQL driver package
go gorm.io/driver/mysql
```

## Software stack
- Go
- Kafka
- Next.js
- MySQL

## Start server and application

``` bash
docker compose up -d --build
```

## Test

After server started

Open [http://localhost:3000](http://localhost:3000) with your browser to test service and application.

## Test managing account by curl command
Add account
```bash
curl --location --request POST 'http://127.0.0.1:4000/openAccount' --header 'Content-Type: application/json' --data-raw '{"AccountHolder": "Test user1","AccountType": 1,"OpeningBalance": 1000000.00}'
```

Deposit
``` bash
curl --location --request POST 'http://127.0.0.1:4000/depositFund' --header 'Content-Type: application/json' --data-raw '{"ID":"Input_ID","Amount":20000}'
```

Withdraw
``` bash
curl --location --request POST 'http://127.0.0.1:4000/withdrawFund' --header 'Content-Type: application/json' --data-raw '{"ID":"Input_ID","Amount":20000}'
```

Close account
``` bash
curl --location --request POST 'http://127.0.0.1:4000/closeAccount' --header 'Content-Type: application/json' --data-raw '{"ID":"Input_ID"}'
```
