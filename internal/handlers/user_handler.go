package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

func (h *Handler) UsersGetHandler(w http.ResponseWriter, r *http.Request) {

	users, err := h.service.User.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func (h *Handler) UsersPostHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdUser, err := h.service.User.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

}

func (h *Handler) UserGetHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.User.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) UserPutHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userUpdated, err := h.service.User.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userUpdated)

}

func (h *Handler) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.User.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
