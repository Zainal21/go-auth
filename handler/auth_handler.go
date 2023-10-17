package handler

import (
	"go-auth/entity"
	"go-auth/repository"
	"go-auth/utils"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request, userRepository *repository.UserRepository) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// fetch user from repository
	user, err := userRepository.FindUserByUserName(username)

	if err != nil || user.Password != password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// generate token
	token, err := utils.GenerateToken(username)

	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token":"` + token + `"}`))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request, userRepository *repository.UserRepository) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var payload entity.User
	payload.Username = username
	payload.Password = password

	err := userRepository.CreateUser(&payload)

	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success":"true"}`))
}
