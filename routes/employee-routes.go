package routes

import (
	"crud/controller"
	"crud/dbconfig"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleEmpRoutes(r *mux.Router, db *dbconfig.DB) {
	empHandeler := controller.NewEmployeeHandler(db)
	r.HandleFunc("/employee", empHandeler.CreateEmployee).Methods(http.MethodPost)
	r.HandleFunc("/employee", empHandeler.GetEmployeesList).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empHandeler.GetEmployeeById).Methods(http.MethodGet)
	r.HandleFunc("/employee/{id}", empHandeler.UpdateEmployee).Methods(http.MethodPut)
	r.HandleFunc("/employee/{id}", empHandeler.DeleteEmployee).Methods(http.MethodDelete)
}
