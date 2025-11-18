package main

import (
	"log"

	"github.com/JiquanZhong/realworld-go/db"
	"github.com/JiquanZhong/realworld-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("Starting server on :8080")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
