package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/dto"
	"onlineShop/internal/utils"
)

// OrdersPostHandler godoc
// @Summary      Создать заказ
// @Description  Создает новый заказ пользователя с указанными товарами
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        order  body      dto.OrderCreateRequest  true  "Данные для создания заказа"
// @Success      201    {object}  dto.OrderResponse
// @Failure      400    {string}  string "Некорректный запрос"
// @Failure      500    {string}  string "Ошибка сервера"
// @Router       /api/orders [post]
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

// OrderGetHandler godoc
// @Summary      Получить заказ по ID
// @Description  Возвращает информацию о заказе по его идентификатору
// @Tags         Orders
// @Produce      json
// @Param        id   path      int  true  "ID заказа"
// @Success      200  {object}  dto.OrderResponse
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/orders/{id} [get]
func (h *Handler) OrderGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	order, err := h.service.Order.GetOrder(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

// OrderPutHandler godoc
// @Summary      Обновить статус заказа
// @Description  Изменяет статус заказа по ID
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        id    path      int                    true  "ID заказа"
// @Param        body  body      dto.OrderUpdateRequest true  "Новый статус заказа"
// @Success      200   {object}  dto.OrderResponse
// @Failure      400   {string}  string "Некорректный запрос"
// @Failure      500   {string}  string "Ошибка сервера"
// @Router       /api/orders/{id} [put]
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

// OrderDeleteHandler godoc
// @Summary      Удалить заказ
// @Description  Удаляет заказ по ID
// @Tags         Orders
// @Param        id   path      int  true  "ID заказа"
// @Success      204  {string}  string "Удалено успешно"
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/orders/{id} [delete]
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

// GetOrdersByUserIDHandler godoc
// @Summary      Получить заказы пользователя
// @Description  Возвращает все заказы, принадлежащие конкретному пользователю
// @Tags         Orders
// @Produce      json
// @Param        id   path      int  true  "ID пользователя"
// @Success      200  {array}   dto.OrderResponse
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/users/{id}/orders [get]
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

// DeleteProductFromOrderHandler godoc
// @Summary      Удалить товар из заказа
// @Description  Удаляет товар из указанного заказа по ID заказа и товара
// @Tags         Orders
// @Param        orderId    path      int  true  "ID заказа"
// @Param        productId  path      int  true  "ID товара"
// @Success      204        {string}  string "Товар удален из заказа"
// @Failure      400        {string}  string "Некорректные параметры"
// @Failure      500        {string}  string "Ошибка сервера"
// @Router       /api/orders/{orderId}/products/{productId} [delete]
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
