package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/12Ndraaa/restapi-jualbeli/config"
	"github.com/12Ndraaa/restapi-jualbeli/routes"
)

func main() {
	//koneksi ke db dlu
	db := config.InitDB()
	defer db.Close()
	routes.RegisterRoutes(db)

	// jlnin srver
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Printf("Server Jalan di Port: %s \n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
