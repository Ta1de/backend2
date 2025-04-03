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

// @Summary      Создать товар
// @Description  Добавляет новый товар в систему
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body  response.CreateProduct  true  "Данные товара"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string  "Ошибка при разборе данных"
// @Failure      500  {object}  map[string]string  "Ошибка при создании товара"
// @Router       /product/create [post]
func (h *Handler) createProduct(c *gin.Context) {
	var productReq response.CreateProduct

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	product := mapper.ToProductModel(productReq)

	id, err := h.services.CreateProduct(c, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось создать товар: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Товар успешно создан",
		"id":      id.String(),
	})
}

// @Summary      Уменьшить количество товара на складе
// @Description  Уменьшает количество указанного товара на складе
// @Tags         products
// @Produce      json
// @Param        id        query  string  true  "UUID товара"
// @Param        quantity  query  int     true  "Количество для уменьшения"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string  "Неверный формат UUID или количества"
// @Failure      404  {object}  map[string]string  "Ошибка при уменьшении товара"
// @Router       /product/updateQuantity [patch]
func (h *Handler) reduceStock(c *gin.Context) {
	productIDParam := c.Query("id")
	productID, err := uuid.Parse(productIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID"})
		return
	}

	quantityParam := c.Query("quantity")
	quantity, err := strconv.Atoi(quantityParam)
	if err != nil || quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Количество товара должно быть положительным числом"})
		return
	}

	err = h.services.ReduceStock(c, productID, quantity)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Количество товара уменьшено"})
}

// @Summary      Получить товар по ID
// @Description  Возвращает информацию о товаре по его UUID
// @Tags         products
// @Produce      json
// @Param        id  path  string  true  "UUID товара"
// @Success      200  {object}  response.ProductResponse
// @Failure      400  {object}  map[string]string  "Некорректный формат ID"
// @Failure      404  {object}  map[string]string  "Ошибка при получении товара"
// @Router       /product/{id} [get]
func (h *Handler) getProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID товара обязателен"})
		return
	}

	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный формат ID"})
		return
	}

	product, err := h.services.GetProductById(c, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении товара: %s", err.Error())})
		return
	}

	productResponse := mapper.ToProductResponse(product)

	c.JSON(http.StatusOK, gin.H{
		"product": productResponse,
	})
}

// @Summary      Получить список товаров
// @Description  Возвращает список всех товаров
// @Tags         products
// @Produce      json
// @Success      200  {array}   response.ProductResponse
// @Failure      404  {object}  map[string]string  "Ошибка при получении товаров"
// @Router       /product/productList [get]
func (h *Handler) getProductList(c *gin.Context) {
	products, err := h.services.GetProductList(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении товара: %s", err.Error())})
		return
	}

	productResponses := make([]response.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = mapper.ToProductResponse(product)
	}

	c.JSON(http.StatusOK, gin.H{
		"products": productResponses,
	})
}

// @Summary      Удалить товар
// @Description  Удаляет товар по его UUID
// @Tags         products
// @Produce      json
// @Param        id  path  string  true  "UUID товара"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string  "Неверный формат UUID"
// @Failure      404  {object}  map[string]string  "Ошибка при удалении товара"
// @Router       /product/delete/{id} [delete]
func (h *Handler) deleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID товара"})
		return
	}

	err = h.services.RemoveProduct(c, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при удалении товара: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Товар успешно удалён"})
}
