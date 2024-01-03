package database

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/storage/mongodb/v2"
	"github.com/joho/godotenv"
)

func connectDb() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	mongoURI := fmt.Sprintf("mongodb+srv://delaricch:%s@product-feedback.l8phg4q.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_PASSWORD"))

	store := mongodb.New(mongodb.Config{
		ConnectionURI: mongoURI,
	})

	er := store.Conn()

	defer func()  {
		if err := store.Close(); er != nil {
			log.Fatalf("Error closing database connection: %s", err)
		}	
	}()

	
}