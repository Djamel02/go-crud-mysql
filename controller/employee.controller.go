package controller

import (
	"crud/dbconfig"
	"crud/models"
	repo "crud/service"
	"crud/service/employee"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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
		respondWithError(w, http.StatusNotFound, err.Error())
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

// File Upload
func fileUpload(r *http.Request) (string, error){
	// Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)


	file, header, err := r.FormFile("picture")
	if err != nil {
		return "", err
	}

	defer file.Close()

	fileExt := strings.Split(header.Filename,".")[1]

	imgName := fmt.Sprintf("uploaded-*.%s", fileExt)

	tempFile, err := ioutil.TempFile("images", imgName)
    if err != nil {
        return "", err
    }

    err = tempFile.Close()
	if err != nil {
        return "", err
    }
	 // read all of the contents of our uploaded file into a
    // byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }
    // write this byte array to our temporary file
    tempFile.Write(fileBytes)
	return tempFile.Name(), nil
}
// Create Employee
func (e *Employee) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	file,err := fileUpload(r)

	req := models.Employee{}
	req.Name =    r.FormValue("name")
	req.Phone =     r.FormValue("phone")
	if err == nil {
		req.Picture =    file
	}

	res, err := e.repo.Create(r.Context(), &req)
	if err != nil {

		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
}

// 

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
