package controller

import (
	"crud/dbconfig"
	"crud/models"
	repo "crud/service"
	"crud/service/employee"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Employee struct {
	repo repo.EmpRepo
}

func NewEmployeeHandler(db *dbconfig.DB) *Employee {
	return &Employee{
		repo: employee.NewEmpRepo(db.SQL),
	}
}

// var views = tmp.Must(tmp.ParseGlob("views/*"))

// Get employees list
func (e *Employee) GetEmployeesList(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.Fetch(r.Context())
	if err != nil {

		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
	// views.ExecuteTemplate(w, "index", res)
}

// Get employee by id
func (e *Employee) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {

		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	res, err := e.repo.GetByID(r.Context(), id)
	if err != nil {

		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
	// views.ExecuteTemplate(w, "index", res)
}

// Create Employee
func (e *Employee) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	req := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	res, err := e.repo.Create(r.Context(), &req)
	if err != nil {

		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
}

// Update employee
func (e *Employee) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	req := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.Update(r.Context(), &req, id)

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, http.StatusAccepted, res)
}

// Delete employee
func (e *Employee) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {

		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.Delete(r.Context(), id)
	if err != nil {

		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, http.StatusOK, res)
}
