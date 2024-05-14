package main

import (
	"log"
	"payment_gateway_mercadopago/routes"
	"payment_gateway_mercadopago/storage"
)

func main() {
	storage.InitializeDB()
	storage.MigrateModels()
	//storage.SeedData()

	r := routes.SetupRoutes()
	log.Println("[--->>>> STARTING SERVER... <<<<---]")
	if err := r.Run(); err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
