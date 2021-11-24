package routehandler

import (
	"crud/controller"
	"crud/dbconfig"
	"crud/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleEmpRoutes(r *mux.Router, db *dbconfig.DB) {
	empHandeler := controller.NewEmployeeHandler(db)
	r.HandleFunc("/employee", middlewares.IsAuthorized(empHandeler.CreateEmployee)).Methods(http.MethodPost)
	r.HandleFunc("/employee", middlewares.IsAuthorized(empHandeler.GetEmployeesList)).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", middlewares.IsAuthorized(empHandeler.GetEmployeeById)).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", middlewares.IsAuthorized(empHandeler.UpdateEmployee)).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}",middlewares.IsAuthorized(empHandeler.DeleteEmployee)).Methods(http.MethodDelete)
}
