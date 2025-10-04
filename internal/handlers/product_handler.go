package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

func (h *Handler) ProductsGetHandler(w http.ResponseWriter, r *http.Request) {

	Products, err := h.service.Product.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Products)

}

func (h *Handler) ProductsPostHandler(w http.ResponseWriter, r *http.Request) {

	var Product models.Product
	if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdProduct, err := h.service.Product.CreateProduct(&Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)

}

func (h *Handler) ProductGetHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Product, err := h.service.Product.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Product)

}

func (h *Handler) ProductPutHandler(w http.ResponseWriter, r *http.Request) {

	var Product models.Product
	if err := json.NewDecoder(r.Body).Decode(&Product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ProductUpdated, err := h.service.Product.UpdateProduct(&Product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ProductUpdated)

}

func (h *Handler) ProductDeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Product.DeleteProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
