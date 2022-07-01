package api

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/kderosha/KongServices/app/db"
	"log"
	"encoding/json"
	"context"
    "github.com/gorilla/mux"
)

type Service struct{
	Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Versions []string `json:"versions" bson:"versions"`
}

type CreateVersion struct{
	Version string `json:"version"`
}

type ServicesResponse struct{
	Services []Service `json:"services"`
}

type VersionsResponse struct{
	Versions []string `json:"versions"`
}

type Api struct{
	dbClient *mongo.Client
}

func AssignHandlers(router *mux.Router) {
    // Return a new router
    // Define routes and handlers.
    log.Println("Registering Handlers");
	servicesApi := &Api{
		dbClient: db.NewDb(),
	}
    router.HandleFunc("/healthz", servicesApi.WelcomeHandler).Methods("GET")
    router.HandleFunc("/services", servicesApi.GetServices).Methods("GET");
    router.HandleFunc("/services", servicesApi.PostServices).Methods("POST");
    router.HandleFunc("/services/{idService}", servicesApi.GetServiceById).Methods("GET");
    router.HandleFunc("/services/{idService}/versions", servicesApi.CreateVersion).Methods("POST");
    router.HandleFunc("/services/{idService}/versions", servicesApi.GetServiceVersions).Methods("GET");
}

// Simple welcome/health handler that can be used to test status of the service.
func (api *Api) WelcomeHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Welcome endpoint");
    fmt.Fprintf(w, "Hello there user");
    return;
}

/// Query the collection of services based on parameters described by the api.yaml for the GET /services endpoint
/// Query parameters
/// 	sort: string, possible options include [asc, desc]
/// 	name: string applied to an open ended regex to search on the name property
func (api *Api) GetServices(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()

	// Define a bson map[string]interface{} to hold the query we will be building to send to the db
	var databaseQuery bson.M = make(bson.M)

	// Get the name query parameter
	nameSearch := query.Get("name")
	sort := query.Get("sort")
	findOptions := options.Find()
	// If the name parameter was included in the request. We add it to the database query
	if nameSearch != "" {
		databaseQuery["name"] = bson.D{{"$regex", ".*" + nameSearch + ".*"}}
	}
	// Based on the sort=1,-1 query parameter evaluate applying a sort to the query
	if sort == "asc" {
		findOptions.SetSort(bson.D{{ "name", 1}})
	} else if sort == "desc"{
		findOptions.SetSort(bson.D{{ "name", -1}})
	}

	collection := api.dbClient.Database("KongServices").Collection("services")
	// Search the mongo collection for our database query
	cursor, err := collection.Find(
		context.TODO(), // TODO: evaluate context of the find operation
		databaseQuery,	// query formed from query parameters on the api
		findOptions) 	// Options considered when searching for documents
	log.Println("searched for collection on ", databaseQuery)

	// Define an array of services
	for cursor.Next(context.TODO()) {
		var service Service
		cursor.Decode(&service)
		log.Println(service)
	}
	var services []Service
	if err = cursor.All(context.TODO(), &services); err != nil {
		panic(err)
	}
	log.Println(services)
	var response *ServicesResponse = &ServicesResponse{
		Services: services,
	}

	// write the array of services to response along with any other information needed
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (api *Api) CreateVersion(w http.ResponseWriter, r *http.Request){
	// Get the service id from the path variable
	params := mux.Vars(r)
	idService := params["idService"]
	log.Println("Received idService: ", idService)
	var createVersion CreateVersion
	err := json.NewDecoder(r.Body).Decode(&createVersion)
	// Validations for the body

	serviceObjectId, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		log.Println("Unable to process id")
		return
	}

	collection := api.dbClient.Database("KongServices").Collection("services")
	if err != nil {
		log.Println("Unable to get collection");
	}

	result, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": serviceObjectId},
		bson.D{
			{"$push", bson.D{{"versions", createVersion.Version}}},
		},
	);
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated ", result.ModifiedCount, " documents")
}

// Handler that searches for a service given the _id for the service document
func (api *Api) GetServiceById(w http.ResponseWriter, r *http.Request){
	// Get the service id from the path variable
	params := mux.Vars(r)
	idService := params["idService"]
	log.Println("Received idService: ", idService)

	serviceObjectId, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		log.Println("Unable to process id")
		return
	}


	// Else we write the service to the response as a json
	service, err := api.getService(serviceObjectId)
	if err != nil {

		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
	return;
}

func (api *Api) GetServiceVersions(w http.ResponseWriter, r *http.Request){
	// Get the service id from the path variable
	params := mux.Vars(r)
	idService := params["idService"]
	log.Println("Received idService: ", idService)

	serviceObjectId, err := primitive.ObjectIDFromHex(idService)
	if err != nil {
		log.Println("Unable to process id")
		return
	}


	// Else we write the service to the response as a json
	service, err := api.getService(serviceObjectId)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
	return;
}

func (api *Api) PostServices(w http.ResponseWriter, r *http.Request){
	// create Service variable
	var service Service

	log.Println("Decoding service")
	// Decode the request body into the service variable based on the json
	err := json.NewDecoder(r.Body).Decode(&service)
	service.Versions = make([]string, 0)

	// do validations on the service structure
	log.Println("Service decoded: ", service)
	if err != nil {
		fmt.Fprintf(w, "error occured: ", err)
		return;
	}

	// store the service in the db, could create a new package for repository layer and use a ORM entity framework.
	log.Println("Inserting document into the db: ", service)
	servicesCollection := api.dbClient.Database("KongServices").Collection("services")
	resultSet, err := servicesCollection.InsertOne(context.TODO(), service)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted document with _id:", resultSet.InsertedID)
	

	// Write that we created a new service to the response.
	// We could write back json
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "Created Service: %+v, with id: %v", service, resultSet.InsertedID)
    return
}

func (api *Api) getService(idService primitive.ObjectID) (*Service, error) {
	// Make query to database in order to retrieve the service given the id.
	collection := api.dbClient.Database("KongServices").Collection("services")

	var returnService *Service
	err := collection.FindOne(context.TODO(), bson.D{{"_id", idService}}).Decode(&returnService)
	if err != nil {
		return nil, err
	}
	return returnService, nil
}