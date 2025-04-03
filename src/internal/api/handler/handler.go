package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"src/internal/service"

	_ "src/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFile("/swagger.json", "./docs/openapi.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.json")))

	apiV1 := router.Group("/api/v1")
	{
		h.initUserRoutes(apiV1)
		h.initSupplierRoutes(apiV1)
		h.initProductRoutes(apiV1)
		h.initImageRoutes(apiV1)
	}

	return router
}

func (h *Handler) initUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.POST("/create", h.createUser)
		user.DELETE("/delete/:id", h.deleteUser)
		user.GET("/users", h.getUsers)
		user.GET("/usersList", h.getUserList)
		user.PUT("/updateAddress/:id", h.updateUserAddress)
	}
}

func (h *Handler) initSupplierRoutes(rg *gin.RouterGroup) {
	supplier := rg.Group("/supplier")
	{
		supplier.POST("/create", h.createSupplier)
		supplier.PUT("/updateAddress/:id", h.updateSupplierAddress)
		supplier.DELETE("/delete/:id", h.deleteSupplier)
		supplier.GET("/supplierList", h.getSupplierList)
		supplier.GET("/:id", h.getSupplier)
	}
}

func (h *Handler) initProductRoutes(rg *gin.RouterGroup) {
	product := rg.Group("/product")
	{
		product.POST("/create", h.createProduct)
		product.PATCH("/updateQuantity", h.reduceStock)
		product.GET("/:id", h.getProduct)
		product.GET("/productList", h.getProductList)
		product.DELETE("/delete/:id", h.deleteProduct)
	}
}

func (h *Handler) initImageRoutes(rg *gin.RouterGroup) {
	image := rg.Group("/image")
	{
		image.POST("/create", h.createImage)
		image.PUT("/updateImage", h.updateImage)
		image.DELETE("/delete/:id", h.deleteImage)
		image.GET("/product/:id", h.getImageByProductId)
		image.GET("/:id", h.getImageById)
	}
}
