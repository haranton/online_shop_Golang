package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/dto"
	"onlineShop/internal/utils"
)

func (h *Handler) OrdersPostHandler(w http.ResponseWriter, r *http.Request) {

	var order dto.OrderCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdOrder, err := h.service.Order.CreateOrder(order.UserID, order.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdOrder)

}

func (h *Handler) OrderGetHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Order, err := h.service.Order.GetOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Order)
}

func (h *Handler) OrderPutHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateReq dto.OrderUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&updateReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedOrder, err := h.service.Order.UpdateStatusOrder(id, updateReq.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedOrder)
}

func (h *Handler) OrderDeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Order.DeleteOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetOrdersByUserIDHandler(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders, err := h.service.Order.GetOrdersByUserID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

func (h *Handler) DeleteProductFromOrderHandler(w http.ResponseWriter, r *http.Request) {

	orderID, err := utils.GetParamIDFromRequest(r, "orderId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	productID, err := utils.GetParamIDFromRequest(r, "productId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Order.DeleteProductFromOrder(orderID, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
