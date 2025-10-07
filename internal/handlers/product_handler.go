package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

// ProductsGetHandler godoc
// @Summary      Получить список товаров
// @Description  Возвращает список всех доступных товаров
// @Tags         Products
// @Produce      json
// @Success      200  {array}   dto.ProductResponse
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/products [get]
func (h *Handler) ProductsGetHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.Product.GetProducts()
	if err != nil {
		h.logger.Error("failed to get products", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

// ProductsPostHandler godoc
// @Summary      Создать товар
// @Description  Добавляет новый товар в систему
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      dto.ProductRequest  true  "Данные нового товара"
// @Success      201      {object}  dto.ProductResponse
// @Failure      400      {string}  string "Некорректные данные"
// @Failure      500      {string}  string "Ошибка сервера"
// @Router       /api/products [post]
func (h *Handler) ProductsPostHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		h.logger.Error("failed to decode product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdProduct, err := h.service.Product.CreateProduct(&product)
	if err != nil {
		h.logger.Error("failed to create product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// ProductGetHandler godoc
// @Summary      Получить товар по ID
// @Description  Возвращает информацию о товаре по его идентификатору
// @Tags         Products
// @Produce      json
// @Param        id   path      int  true  "ID товара"
// @Success      200  {object}  dto.ProductResponse
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/products/{id} [get]
func (h *Handler) ProductGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		h.logger.Error("failed to get product id from request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.Product.GetProduct(id)
	if err != nil {
		h.logger.Error("failed to get product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// ProductPutHandler godoc
// @Summary      Обновить товар
// @Description  Обновляет информацию о товаре по ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id       path      int               true  "ID товара"
// @Param        product  body      dto.ProductRequest true  "Обновленные данные товара"
// @Success      200      {object}  dto.ProductResponse
// @Failure      400      {string}  string "Некорректные данные"
// @Failure      500      {string}  string "Ошибка сервера"
// @Router       /api/products/{id} [put]
func (h *Handler) ProductPutHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		h.logger.Error("failed to decode product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	productUpdated, err := h.service.Product.UpdateProduct(&product)
	if err != nil {
		h.logger.Error("failed to update product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productUpdated)
}

// ProductDeleteHandler godoc
// @Summary      Удалить товар
// @Description  Удаляет товар по ID
// @Tags         Products
// @Param        id   path      int  true  "ID товара"
// @Success      204  {string}  string "Товар успешно удален"
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/products/{id} [delete]
func (h *Handler) ProductDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		h.logger.Error("failed to get product id from request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Product.DeleteProduct(id)
	if err != nil {
		h.logger.Error("failed to delete product", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
