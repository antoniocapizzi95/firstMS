package main

import (
	"context"
	"firstMS/repository"
	"firstMS/repository/kafkarepo"
	"firstMS/repository/mongorepo"
	"firstMS/reqhandlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	exposeOnAddress := os.Getenv("EXPOSE_ON_ADDRESS")
	exposeOnPort := os.Getenv("EXPOSE_ON_PORT")
	kafkaAddress := os.Getenv("KAFKA_ADDRESS")
	kafkaPort := os.Getenv("KAFKA_PORT")
	mongoAddress := os.Getenv("MONGO_ADDRESS")
	mongoPort := os.Getenv("MONGO_PORT")

	abDb := getDbRepository(mongoAddress, mongoPort)
	abQueue := getQueueRepository(kafkaAddress, kafkaPort)
	repository.InitModule(abDb, abQueue)

	r := chi.NewRouter()
	r.Get("/", reqhandlers.GetAddressBook)
	r.Post("/", reqhandlers.AddPerson)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", exposeOnAddress, exposeOnPort), r)
	if err != nil {
		return
	}
}

func getDbRepository(address string, port string) *repository.AddressBookRepo {
	address = fmt.Sprintf("mongodb://%s:%s", address, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(address))
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()

	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	db := client.Database("testms")
	abDb := mongorepo.GetAddressBookMongoRepo(db.Collection("AddressBook"))
	return &abDb
}

func getQueueRepository(address string, port string) *repository.AddressBookRepo {
	address = fmt.Sprintf("%s:%s", address, port)
	abQueue := kafkarepo.GetAddressBookKafkaRepo(address, "add-person", 0)
	return &abQueue
}
