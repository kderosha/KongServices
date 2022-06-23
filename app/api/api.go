package api

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"encoding/json"
	"context"
)

type CreateServiceModel struct{
	name string
	description string
}

type Service struct{
	Id string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Versions []string `json:"versions" bson:"versions"`
	VersionCount int32
}

type Api struct{
	db *mongo.Client
}

func NewApi(db *mongo.Client)  *Api{
	return &Api{
		db: db,
	}
}

func (api *Api) WelcomeHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Welcome endpoint");

    fmt.Fprintf(w, "Hello there user");
    return;
}

func (api *Api) GetServices(w http.ResponseWriter, r *http.Request){
	// marshal the service 
    return
}

func (api *Api) PostServices(w http.ResponseWriter, r *http.Request){

	// create Service variable
	var service CreateServiceModel

	log.Println("Decoding service")
	// Decode the request body into the service variable based on the json
	err := json.NewDecoder(r.Body).Decode(&service)
	log.Println("Service decoded: %+v", service)
	if err != nil {
		fmt.Fprintf(w, "error occured: %+v", err)
		return;
	}

	// store the service in the db, could create a new package for repository layer and use a ORM entity framework.
	log.Println("Inserting document into the db: %+v", service)
	servicesCollection := api.db.Database("KongServices").Collection("services")
	resultSet, err := servicesCollection.InsertOne(context.TODO(), service)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted document with _id: %v", resultSet.InsertedID)
	

	// Write that we created a new service to the response.
	// We could write back json
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "Created Service: %+v, with id: %v", service, resultSet.InsertedID)
    return
}