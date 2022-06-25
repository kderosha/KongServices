package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "github.com/kderosha/KongServices/app/api"
)

func main(){
    log.Println("Starting service");
    log.Println("Register some observability mechanisms")

    router := mux.NewRouter();
    router.Use(authMiddleware)
    // Configure CORS for our front end application.

    // Register the new api with the database.
    log.Println("Registering API")
    api.AssignHandlers(router)

    log.Println("Starting server")
    srv := &http.Server{
        Handler: router,
        Addr: ":8000",
    }
    log.Fatal(srv.ListenAndServe());
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        log.Println("Perform authentication using OAuth2 and OIDC")
        next.ServeHTTP(w, r)
    })
}