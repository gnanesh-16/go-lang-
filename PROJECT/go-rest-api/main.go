package main

import (
    "log"
    "net/http"
    "go-rest-api/config"
    "go-rest-api/routes"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    config.ConnectDatabase()

    routes.RegisterRoutes(router)

    log.Fatal(http.ListenAndServe(":8080", router))
}
