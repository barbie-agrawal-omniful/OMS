package main

import (
	"fmt"
	"log"
	"time"

	"github.com/omniful/go_commons/http"
)

func main() {
	server := http.InitializeServer(":8081", 10*time.Second, 10*time.Second, 70*time.Second)
	fmt.Println("Starting server...")
	if err := server.StartServer("OMS Service Started"); err != nil {
		log.Fatalf("Error starting the server")
	}
}
