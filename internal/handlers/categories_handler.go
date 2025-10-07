package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	_ "onlineShop/internal/dto"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

// CategoriesGetHandler godoc
// @Summary      Получить список категорий
// @Description  Возвращает список всех категорий товаров
// @Tags         Categories
// @Produce      json
// @Success      200  {array}   dto.CategoryResponse
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/categories [get]
func (h *Handler) CategoriesGetHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.Category.GetCategories()
	if err != nil {
		h.logger.Error("failed to get categories", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// CategoriesPostHandler godoc
// @Summary      Создать категорию
// @Description  Добавляет новую категорию товаров
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        category  body      dto.CategoryRequest  true  "Данные категории"
// @Success      201       {object}  dto.CategoryResponse
// @Failure      400       {string}  string "Некорректные данные"
// @Failure      500       {string}  string "Ошибка сервера"
// @Router       /api/categories [post]
func (h *Handler) CategoriesPostHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Error("failed to decode category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createdCategory, err := h.service.Category.CreateCategory(&category)
	if err != nil {
		h.logger.Error("failed to create category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCategory)
}

// CategoryGetHandler godoc
// @Summary      Получить категорию по ID
// @Description  Возвращает категорию по её идентификатору
// @Tags         Categories
// @Produce      json
// @Param        id   path      int  true  "ID категории"
// @Success      200  {object}  dto.CategoryResponse
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/categories/{id} [get]
func (h *Handler) CategoryGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		h.logger.Error("failed to get category id from request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	category, err := h.service.Category.GetCategory(id)
	if err != nil {
		h.logger.Error("failed to get category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

// CategoryPutHandler godoc
// @Summary      Обновить категорию
// @Description  Обновляет данные категории по ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        id         path      int                 true  "ID категории"
// @Param        category   body      dto.CategoryRequest  true  "Обновленные данные категории"
// @Success      200        {object}  dto.CategoryResponse
// @Failure      400        {string}  string "Некорректные данные"
// @Failure      500        {string}  string "Ошибка сервера"
// @Router       /api/categories/{id} [put]
func (h *Handler) CategoryPutHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		h.logger.Error("failed to decode category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoryUpdated, err := h.service.Category.UpdateCategory(&category)
	if err != nil {
		h.logger.Error("failed to update category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoryUpdated)
}

// CategoryDeleteHandler godoc
// @Summary      Удалить категорию
// @Description  Удаляет категорию по ID
// @Tags         Categories
// @Param        id   path      int  true  "ID категории"
// @Success      204  {string}  string "Категория успешно удалена"
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/categories/{id} [delete]
func (h *Handler) CategoryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamIDFromRequest(r, "id")
	if err != nil {
		h.logger.Error("failed to get category id from request", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Category.DeleteCategory(id)
	if err != nil {
		h.logger.Error("failed to delete category", slog.String("error", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
