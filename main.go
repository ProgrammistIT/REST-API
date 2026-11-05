package main

import (
	"fmt"
	"net/http"

	"threat-api/internal/database"
	"threat-api/internal/handler"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("POST /info", handler.CreateThreat(db))
	http.HandleFunc("GET /info/{md5}", handler.GetThreat(db))
	http.HandleFunc("PUT /info/{md5}", handler.UpdateThreat(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}
