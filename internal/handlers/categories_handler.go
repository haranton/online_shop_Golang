package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

func (h *Handler) CategoriesGetHandler(w http.ResponseWriter, r *http.Request) {

	Categories, err := h.service.Category.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Categories)

}

func (h *Handler) CategoriesPostHandler(w http.ResponseWriter, r *http.Request) {

	var Category models.Category
	if err := json.NewDecoder(r.Body).Decode(&Category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdCategory, err := h.service.Category.CreateCategory(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)

}

func (h *Handler) CategoryGetHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Category, err := h.service.Category.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Category)

}

func (h *Handler) CategoryPutHandler(w http.ResponseWriter, r *http.Request) {

	var Category models.Category
	if err := json.NewDecoder(r.Body).Decode(&Category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	CategoryUpdated, err := h.service.Category.UpdateCategory(&Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CategoryUpdated)

}

func (h *Handler) CategoryDeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Category.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
