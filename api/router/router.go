package router

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/products/", controller.GetAllProduct)
		api.GET("/Product/:nama_product", controller.FindProduct)
		api.POST("/CreateProduct", controller.CreateProduct)
		api.POST("/DeleteProduct/:nama_product", controller.DeleteProduct)
	}
}