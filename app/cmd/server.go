package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "github.com/kderosha/KongServices/app/db"
    "github.com/kderosha/KongServices/app/api"
)

func main(){
    // Regester observability and logging framework
    log.Println("Starting service");
    log.Println("Registering database");
    db := db.NewDb();

    log.Println("Registering API")
    servicesApi := api.NewApi(db)

    log.Println("Registering Handlers");
    router := mux.NewRouter();
    router.HandleFunc("/", servicesApi.WelcomeHandler).Methods("GET")
    router.HandleFunc("/services", servicesApi.GetServices).Methods("GET");
    router.HandleFunc("/services", servicesApi.PostServices).Methods("POST");
    log.Println("Starting service")
    srv := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:8000",
    }
    log.Fatal(srv.ListenAndServe());
}
