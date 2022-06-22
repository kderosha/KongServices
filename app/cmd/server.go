package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "log"
    // "encoding/json"
)

func main(){
    // Regester observability and logging framework
    fmt.Println("Starting service");
    fmt.Println("Registering Handlers");
    router := mux.NewRouter();
    router.HandleFunc("/", WelcomeHandler)
    router.HandleFunc("/services", GetServices).Methods("GET");
    router.HandleFunc("/services", PostServices).Methods("POST");
    srv := &http.Server{
        Handler: router,
        Addr:    "127.0.0.1:8080",
    }
    log.Fatal(srv.ListenAndServe());
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello there user");
    return;
}

func GetServices(w http.ResponseWriter, r *http.Request){
    return
}

func PostServices(w http.ResponseWriter, r *http.Request){
    return
}
