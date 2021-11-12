package controller

import (
	"crud/dbconfig"
	"crud/models"
	"crud/service/employee"
	"encoding/json"
	"net/http"
)

// new employee handlerr ...
func newEmployeeHandler(db *dbconfig.DB) *models.Employee {
	return &models.Employee{
		repo: employee.NewSQLEmpRepo(db.SQL),
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
