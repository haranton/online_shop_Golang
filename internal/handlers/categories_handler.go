package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

// CategoriesGetHandler godoc
// @Summary      Получить список категорий
// @Description  Возвращает список всех категорий товаров
// @Tags         Categories
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/categories [get]
func (h *Handler) CategoriesGetHandler(w http.ResponseWriter, r *http.Request) {
	Categories, err := h.service.Category.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(Categories)
}

// CategoriesPostHandler godoc
// @Summary      Создать категорию
// @Description  Добавляет новую категорию товаров
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Данные категории"
// @Success      201       {object}  models.Category
// @Failure      400       {string}  string "Некорректные данные"
// @Failure      500       {string}  string "Ошибка сервера"
// @Router       /api/categories [post]
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

// CategoryGetHandler godoc
// @Summary      Получить категорию по ID
// @Description  Возвращает категорию по её идентификатору
// @Tags         Categories
// @Produce      json
// @Param        id   path      int  true  "ID категории"
// @Success      200  {object}  models.Category
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/categories/{id} [get]
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

// CategoryPutHandler godoc
// @Summary      Обновить категорию
// @Description  Обновляет данные категории по ID
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Обновленные данные категории"
// @Success      200       {object}  models.Category
// @Failure      400       {string}  string "Некорректные данные"
// @Failure      500       {string}  string "Ошибка сервера"
// @Router       /api/categories/{id} [put]
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Category.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
