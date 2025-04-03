package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"src/internal/api/response"
	"src/internal/middleware/mapper"
	"strconv"
)

// @Summary Создание пользователя
// @Description Создает нового пользователя на основе переданных данных
// @Tags users
// @Accept json
// @Produce json
// @Param user body response.CreateUser true "Данные нового пользователя"
// @Success 201 {object} map[string]string "Пользователь успешно создан"
// @Failure 400 {object} map[string]string "Ошибка в данных"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /user/create [post]
func (h *Handler) createUser(c *gin.Context) {
	var userReq response.CreateUser

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	address := mapper.ToAddressModel(userReq.Address)
	user := mapper.ToUserModel(userReq)

	id, err := h.services.AddUser(c, user, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось создать пользователя: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан",
		"id":      id.String(),
	})
}

// @Summary Удаление пользователя
// @Description Удаляет пользователя по UUID
// @Tags users
// @Param id path string true "UUID пользователя"
// @Success 200 {object} map[string]string "Пользователь успешно удалён"
// @Failure 400 {object} map[string]string "Неверный формат UUID"
// @Failure 404 {object} map[string]string "Ошибка при удалении пользователя"
// @Router /user/delete/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID пользователя"})
		return
	}

	err = h.services.RemoveUser(c, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при удалении пользователя: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удалён"})
}

// @Summary Получение списка пользователей по имени и фамилии
// @Description Возвращает список пользователей, отфильтрованных по имени и фамилии
// @Tags users
// @Param name query string true "Имя пользователя"
// @Param surname query string true "Фамилия пользователя"
// @Success 200 {object} map[string][]response.UserResponse "Список пользователей"
// @Failure 400 {object} map[string]string "Ошибка в параметрах запроса"
// @Failure 404 {object} map[string]string "Ошибка сервера"
// @Router /user/users [get]
func (h *Handler) getUsers(c *gin.Context) {
	name := c.Query("name")
	surname := c.Query("surname")

	if name == "" || surname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя и фамилия обязательны"})
		return
	}

	users, err := h.services.GetUsers(c, name, surname)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении пользователя: %s", err.Error())})
		return
	}

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = mapper.ToUserResponse(user)
	}

	c.JSON(http.StatusOK, gin.H{"users": userResponses})
}

// @Summary Получение списка пользователей с пагинацией
// @Description Возвращает список пользователей с возможностью пагинации
// @Tags users
// @Param limit query int false "Количество пользователей на странице (по умолчанию 20)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {object} map[string]interface{} "Список пользователей и флаг has_more"
// @Failure 400 {object} map[string]string "Ошибка в параметрах запроса"
// @Failure 404 {object} map[string]string "Ошибка сервера"
// @Router /user/usersList [get]
func (h *Handler) getUserList(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, offset := 20, 0
	var err error

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'limit' должен быть целым числом >= 0"})
			return
		}
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'offset' должен быть целым числом >= 0"})
			return
		}
	}

	users, err := h.services.GetUsersList(c, limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении пользователя: %s", err.Error())})
		return
	}

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = mapper.ToUserResponse(user)
	}

	hasMore := len(users) == limit

	c.JSON(http.StatusOK, gin.H{
		"users":    userResponses,
		"has_more": hasMore,
	})
}

// @Summary Обновление адреса пользователя
// @Description Изменяет адрес пользователя по UUID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "UUID пользователя"
// @Param address body response.CreateUpdateAddress true "Новый адрес пользователя"
// @Success 200 {object} map[string]string "Адрес успешно изменен"
// @Failure 400 {object} map[string]string "Ошибка в параметрах запроса"
// @Failure 404 {object} map[string]string "Ошибка сервера"
// @Router /user/updateAddress/{id} [put]
func (h *Handler) updateUserAddress(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID пользователя"})
		return
	}

	var Address response.CreateUpdateAddress

	if err := c.ShouldBindJSON(&Address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	address := mapper.ToAddressModel(Address)

	err = h.services.UpdateUserAddress(c, userID, address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при удалении пользователя: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Адрес успешно изменен"})
}
