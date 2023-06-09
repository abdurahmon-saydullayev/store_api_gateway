package api

import (
	_ "Projects/store/store_api_gateway/api/docs"
	"Projects/store/store_api_gateway/api/handlers"
	"Projects/store/store_api_gateway/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	//user
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetUserByID)
	r.GET("/user", h.GetUserList)
	r.PUT("user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	r.PATCH("/user/:id", h.UpdatePatchUser)
	//order
	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOrderByID)

	// otp
	r.POST("/check", h.CreateUserOTP)
	r.GET("/check", h.VerifyUserOTP)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
