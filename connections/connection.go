package connections

import (
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connection(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGO_URI")


	// Initialize MongoDB connection using mgm
	err = mgm.SetDefaultConfig(nil, "my_database", options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB successfully!")
}