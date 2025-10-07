package handlers

import (
	"encoding/json"
	"net/http"
	"onlineShop/internal/models"
	"onlineShop/internal/utils"
)

// UsersGetHandler godoc
// @Summary      Получить список пользователей
// @Description  Возвращает список всех зарегистрированных пользователей
// @Tags         Users
// @Produce      json
// @Success      200  {array}   dto.UserResponse
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/users [get]
func (h *Handler) UsersGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.User.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

// UsersPostHandler godoc
// @Summary      Создать пользователя
// @Description  Регистрирует нового пользователя в системе
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserRequest  true  "Данные пользователя"
// @Success      201   {object}  dto.UserResponse
// @Failure      400   {string}  string "Некорректные данные"
// @Failure      500   {string}  string "Ошибка сервера"
// @Router       /api/users [post]
func (h *Handler) UsersPostHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

// UserGetHandler godoc
// @Summary      Получить пользователя по ID
// @Description  Возвращает данные пользователя по его идентификатору
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "ID пользователя"
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/users/{id} [get]
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

// UserPutHandler godoc
// @Summary      Обновить данные пользователя
// @Description  Обновляет информацию о пользователе по его ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "ID пользователя"
// @Param        user  body      dto.UserRequest  true  "Обновленные данные пользователя"
// @Success      200   {object}  dto.UserResponse
// @Failure      400   {string}  string "Некорректные данные"
// @Failure      500   {string}  string "Ошибка сервера"
// @Router       /api/users/{id} [put]
func (h *Handler) UserPutHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

// UserDeleteHandler godoc
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по его ID
// @Tags         Users
// @Param        id   path      int  true  "ID пользователя"
// @Success      204  {string}  string "Пользователь успешно удален"
// @Failure      400  {string}  string "Некорректный ID"
// @Failure      500  {string}  string "Ошибка сервера"
// @Router       /api/users/{id} [delete]
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
	w.WriteHeader(http.StatusNoContent)
}
