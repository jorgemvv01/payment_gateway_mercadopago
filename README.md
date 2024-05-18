# Mercado Pago - Payment Gateway / Split payments
A simple Go-API to implement split payments in Mercado Pago.

This Go API was created with Go using the Gin framework, the GORM ORM and PostgreSQL for relational database management.

![UI](https://github.com/jorgemvv01/payment_gateway_mercadopago/blob/main/mp_payment_gateway.gif)

## Installation & Run
**Step 1:**

Download or clone this repo by using the link below:
```
https://github.com/jorgemvv01/payment_gateway_mercadopago
```

**Step 2:**

Create databases:
```sql
CREATE DATABASE mp_gateway;
CREATE DATABASE mp_gateway_test;
```
Configures the database connection with the environment variables on [storage.go](https://github.com/jorgemvv01/go-api/tree/master/storage/storage.go):
```go
func InitializeDB() {
    if DB == nil {
        var err error
        dsn := os.Getenv("DATABASE_URL")
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            panic("failed to connect database")
        }
    }
}
```
```
DATABASE_URL=postgresql://YOUR_USER:YOUR_PASSWORD@HOST:PORT/mp_gateway
DATABASE_URL=postgresql://YOUR_USER:YOUR_PASSWORD@HOST:PORT/mp_gateway_test
```
**Step 3:**

Build:
```bash
cd payment_gateway_mercadopago
go build
```

**Step 4:**

Run all test:
```bash
go test ./...
```

**Step 5:**

Run project:
```bash
go run main.go
```



## API documentation
Go to:
```
your_host/api/docs/index.html
```

## Structure
```
├── adaptar
├───── http
├── docs
├── domain
├───── interfaces
├───── models
├───── services
├── infrastructure
├───── repository
├── mocks
├── routes
├── storage
└── main.go
```
