package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"src/internal/api/response"
	"src/internal/middleware/mapper"
)

// @Summary      Создать изображение
// @Description  Создаёт новое изображение и привязывает его к продукту
// @Tags         images
// @Accept       json
// @Produce      json
// @Param        request body response.UploadUpdateImage true "Данные изображения"
// @Success      201 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /image/create [post]
func (h *Handler) createImage(c *gin.Context) {
	var imageReq response.UploadUpdateImage

	if err := c.ShouldBindJSON(&imageReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	image := mapper.ToImageModel(imageReq)
	productID, err := uuid.Parse(imageReq.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID продукта"})
		return
	}

	id, err := h.services.CreateImage(c, image, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Не удалось создать изображение: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Изображение успешно создано",
		"id":      id.String(),
	})
}

// @Summary      Обновить изображение
// @Description  Обновляет данные изображения по его ID
// @Tags         images
// @Accept       json
// @Produce      json
// @Param        request body response.UploadUpdateImage true "Обновленные данные изображения"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /image/updateImage [put]
func (h *Handler) updateImage(c *gin.Context) {
	var imageReq response.UploadUpdateImage

	if err := c.ShouldBindJSON(&imageReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Ошибка при разборе данных: %s", err.Error())})
		return
	}

	image := mapper.ToImageModel(imageReq)
	imageID, err := uuid.Parse(imageReq.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID продукта"})
		return
	}

	err = h.services.UpdateImage(c, image, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Не удалось создать изображение: %s", err.Error())})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Изображение успешно изменино",
	})
}

// @Summary      Удалить изображение
// @Description  Удаляет изображение по ID
// @Tags         images
// @Produce      json
// @Param        id path string true "UUID изображения"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /image/delete/{id} [delete]
func (h *Handler) deleteImage(c *gin.Context) {
	imageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID изображения"})
		return
	}

	err = h.services.DeleteImage(c, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Не удалось удалить изображение: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Изображение успешно удалено"})
}

// @Summary      Получить изображение по ID продукта
// @Description  Возвращает изображение по UUID продукта
// @Tags         images
// @Produce      octet-stream
// @Param        id path string true "UUID продукта"
// @Success      200 {file} binary
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /image/product/{id} [get]
func (h *Handler) getImageByProductId(c *gin.Context) {
	productID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID продукта"})
		return
	}

	image, err := h.services.GetImageByProductId(c, productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении изображения: %s", err.Error())})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, image.ID))

	c.Data(http.StatusOK, "application/octet-stream", image.Image)
}

// @Summary      Получить изображение по его ID
// @Description  Возвращает изображение по UUID
// @Tags         images
// @Produce      octet-stream
// @Param        id path string true "UUID изображения"
// @Success      200 {file} binary
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /image/{id} [get]
func (h *Handler) getImageById(c *gin.Context) {
	imageID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат UUID изображения"})
		return
	}

	image, err := h.services.GetImageById(c, imageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Ошибка при получении изображения: %s", err.Error())})
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.png"`, imageID))

	c.Data(http.StatusOK, "application/octet-stream", image.Image)
}
