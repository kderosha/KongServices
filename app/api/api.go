package api

import (
	"net/http"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type Api struct{
	db *mongo.Client
}

func NewApi(db *mongo.Client)  *Api{
	return &Api{
		db: db,
	}
}

func (api *Api) WelcomeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello there user");
    return;
}

func (api *Api) GetServices(w http.ResponseWriter, r *http.Request){
    return
}

func (api *Api) PostServices(w http.ResponseWriter, r *http.Request){
    return
}