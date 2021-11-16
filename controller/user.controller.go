package controller

import (
	"crud/dbconfig"
	"crud/models"
	repo "crud/service"
	"crud/service/user"
	"encoding/json"
	"net/http"
)

type User struct {
	repo repo.UserRepo
}

func NewUserHandler(db *dbconfig.DB) *User {
	return &User{
		repo: user.NewUserRepo(db.SQL),
	}
}

func (u *User) RegisterUser(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := u.repo.Register(r.Context(), &req)
	if err != nil {

		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, 200, res)
}
func (u *User) Singin(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := u.repo.Login(r.Context(), req.Username, req.Password)
	if err != nil {

		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, res)
}
