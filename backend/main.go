package main

import (
	"fmt"
	"log"

	"encoding/json"
	"net/http"

	"igracing/database"
	"igracing/models"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	response := Response{Message: "Hello from Go!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	database.InitDB()

	http.HandleFunc("/api/helloworld", handler)
	http.ListenAndServe(":8080", nil)

	// Migrar o esquema
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Database migration completed!")

	// Exemplo de criação de um novo usuário
	newUser := models.User{Name: "John Doe", Email: "john.doe@example.com"}
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		log.Fatal("Failed to create user: ", result.Error)
	}
	fmt.Println("New user created:", newUser)
}
