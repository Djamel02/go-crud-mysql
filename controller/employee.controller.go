package controller

import (
	"crud/dbconfig"
	"crud/models"
	repo "crud/service"
	"crud/service/employee"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee ...
type Employee struct {
	repo repo.EmpRepo
}

// new employee handlerr ...
func newEmployeeHandler(db *dbconfig.DB) *Employee {
	return &Employee{
		repo: employee.NewEmpRepo(db.SQL),
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}

// Get employee by id
func (e *Employee) getEmployeeById(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Errorf("Error While processing Request", err)
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	res, err := e.repo.GetByID(r.Context(), id)
	if err != nil {
		fmt.Errorf("Error While processing Request", err)
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
}

// Create Employee
func (e *Employee) createEmployee(w http.ResponseWriter, r *http.Request) {
	req := models.Employee{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		fmt.Errorf("Error While processing Request", err)
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	res, err := e.repo.Create(r.Context(), &req)
	if err != nil {
		fmt.Errorf("Error While processing Request", err)
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
}
