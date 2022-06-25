package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCreateNewService(t *testing.T){
	req := httptest.NewRequest(http.MethodPost, )
	w := http.NewRecorder()
	// connect to the test database that should be created in CI
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:test@localhost:27017"))
	if err != nil{
		panic(err)
	}
	return client;

	api := &Api{
		dbClient: client,
	}

	api.PostServices(w, req)
	// Assert the response is good

}

