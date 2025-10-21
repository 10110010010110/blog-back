package router

import (
	"blog-back/JWT"
	"blog-back/controller"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/user/register", controller.Register)
	api.POST("/user/login", controller.Login)
	api.PUT("/user/edituser/:id", JWT.JWTAuthMiddleware(), controller.Edituser)
	api.POST("/post/createpost", JWT.JWTAuthMiddleware(), controller.Createpost)
	api.GET("/post/getposts", controller.Getposts)
	api.PUT("post/editpost/:id", JWT.JWTAuthMiddleware(), controller.Editpost)
	api.DELETE("post/deletepost", JWT.JWTAuthMiddleware(), controller.Deletepost)
	api.GET("/post/getpostnumber", controller.Getpostnumber)
	r.Run(":8081")

}
