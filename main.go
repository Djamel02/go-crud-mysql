package main

import (
	"crud/dbconfig"
	"crud/routehandler"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db, err := dbconfig.Connect()
	if err != nil {
		panic(err)
	}
	routehandler.HandleEmpRoutes(router, db)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", dbconfig.GetEnvironmentVars("PORT")), router))

}
