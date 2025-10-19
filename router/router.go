package router

import (
	"blog-back/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/user/register", controller.Register)
	api.POST("/user/login", controller.Login)
	api.PUT("/user/edituser", controller.Edituser)
	api.POST("/post/createpost", controller.Createpost)
	api.GET("/post/getposts", controller.Getposts)
	api.PUT("post/editpost", controller.Editpost)
	api.DELETE("post/deletepost", controller.Deletepost)

	r.Run(":8081")

}
