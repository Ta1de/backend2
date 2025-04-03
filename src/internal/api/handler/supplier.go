package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"src/internal/api/response"
	"src/internal/middleware/mapper"
)

// @Summary Создать поставщика
// @Description Создает нового поставщика с указанным адресом
// @Tags suppliers
// @Accept json
// @Produce json
// @Param supplier body response.CreateSupplier true "Данные нового поставщика"
// @Success 201 {object} map[string]string "Поставщик успешно создан"
// @Failure 400 {object} map[string]string "Ошибка в данных"
// @Failure 500 {object} map[string]string "Не удалось создать поставщика"
// @Router /supplier/create [post]
func (h *Handler) createSupplier(c *gin.Context) {
	var supplierReq response.CreateSupplier

	if err := c.ShouldBindJSON(&supplierReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	address := mapper.ToAddressModel(supplierReq.Address)
	supplier := mapper.ToSupplierModel(supplierReq)

	id, err := h.services.AddSupplier(c, supplier, address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось создать пользователя: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Поставщик успешно создан",
		"id":      id.String(),
	})
}

// @Summary Обновить адрес поставщика
// @Description Обновляет адрес поставщика по его ID
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "UUID поставщика"
// @Param address body response.CreateUpdateAddress true "Новый адрес поставщика"
// @Success 200 {object} map[string]string "Адрес успешно изменен"
// @Failure 400 {object} map[string]string "Ошибка в данных или некорректный UUID"
// @Failure 404 {object} map[string]string "Ошибка при обновлении адреса"
// @Router /supplier/updateAddress/{id} [put]
func (h *Handler) updateSupplierAddress(c *gin.Context) {
	supplierIDStr := c.Param("id")
	supplierID, err := uuid.Parse(supplierIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID поставщика"})
		return
	}

	var Address response.CreateUpdateAddress

	if err := c.ShouldBindJSON(&Address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	address := mapper.ToAddressModel(Address)

	err = h.services.UpdateSupplierAddress(c, supplierID, address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при удалении поставщика: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Адрес успешно изменен"})
}

// @Summary Удалить поставщика
// @Description Удаляет поставщика по его ID
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "UUID поставщика"
// @Success 200 {object} map[string]string "Поставщик успешно удалён"
// @Failure 400 {object} map[string]string "Некорректный UUID поставщика"
// @Failure 404 {object} map[string]string "Ошибка при удалении поставщика"
// @Router /supplier/delete/{id} [delete]
func (h *Handler) deleteSupplier(c *gin.Context) {
	supplierIDStr := c.Param("id")
	supplierID, err := uuid.Parse(supplierIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID поставщика"})
		return
	}

	err = h.services.RemoveSupplier(c, supplierID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при удалении поставщика: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Поставщик успешно удалён"})
}

// @Summary Получить список поставщиков
// @Description Возвращает список всех поставщиков
// @Tags suppliers
// @Accept json
// @Produce json
// @Success 200 {object} map[string][]response.SupplierResponse "Список поставщиков"
// @Failure 404 {object} map[string]string "Ошибка при получении списка"
// @Router /supplier/supplierList [get]
func (h *Handler) getSupplierList(c *gin.Context) {
	supliers, err := h.services.GetSuppliersList(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении пользователя: %s", err.Error())})
		return
	}

	suplierResponses := make([]response.SupplierResponse, len(supliers))
	for i, suplier := range supliers {
		suplierResponses[i] = mapper.ToSupplierResponse(suplier)
	}

	c.JSON(http.StatusOK, gin.H{
		"supliers": suplierResponses,
	})
}

// @Summary Получить поставщика
// @Description Возвращает данные поставщика по его ID
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "UUID поставщика"
// @Success 200 {object} response.SupplierResponse "Данные поставщика"
// @Failure 400 {object} map[string]string "Некорректный UUID или отсутствует ID"
// @Failure 404 {object} map[string]string "Ошибка при получении поставщика"
// @Router /supplier/{id} [get]
func (h *Handler) getSupplier(c *gin.Context) {
	supplierIDStr := c.Param("id")
	if supplierIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID поставщика обязателен"})
		return
	}

	supplierID, err := uuid.Parse(supplierIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат ID"})
		return
	}

	supplier, err := h.services.GetSupplierByID(c, supplierID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении поставщика: %s", err.Error())})
		return
	}

	supplierResponse := mapper.ToSupplierResponse(supplier)

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplierResponse,
	})
}
