package main

import (
	"log"
	_ "payment_gateway_mercadopago/docs"
	"payment_gateway_mercadopago/routes"
	"payment_gateway_mercadopago/storage"
)

// @title Mercado Pago - Payment Gateway / Split payments
// @version 1.0
// @description A simple Go-API to implement split payments in Mercado Pago. \n GitHub Repository: https://github.com/jorgemvv01/payment_gateway_mercadopago

// @contact.name   Jorge Mario Villarreal Vargas.
// @contact.url    https://www.linkedin.com/in/jorgemariovillarreal/
// @contact.email  jorgemvv01@gmail.com

// @BasePath /api
func main() {
	storage.InitializeDB()
	storage.MigrateModels()
	storage.SeedData()

	r := routes.SetupRoutes()
	log.Println("[--->>>> STARTING SERVER... <<<<---]")
	if err := r.Run(); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
