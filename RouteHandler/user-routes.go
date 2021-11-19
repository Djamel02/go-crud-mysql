package routehandler

import (
	"crud/controller"
	"crud/dbconfig"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleUserRoutes(r *mux.Router, db *dbconfig.DB) {
	userHandler := controller.NewUserHandler(db)
	r.HandleFunc("/register", userHandler.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.Singin).Methods(http.MethodPost)
	r.HandleFunc("/users", userHandler.GetUsersList).Methods(http.MethodGet)
}
