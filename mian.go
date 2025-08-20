package main

import (
	"log"
	"vote-backend/config"
	"vote-backend/routes"
)

func main() {
	r := routes.SetupRouter()
	port := config.GetPort()
	log.Printf("Server listening on :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
