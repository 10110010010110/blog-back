package main

import (
	"blog-back/dao"
	"blog-back/models"
	"blog-back/router"
)

func main() {
	dao.Initmysql()
	dao.DB.AutoMigrate(&models.User{})
	dao.DB.AutoMigrate(&models.Post{})
	router.Init()

}
