package controller

import (
	"blog-back/JWT"
	"blog-back/dao"
	"blog-back/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := dao.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": " 注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := dao.DB.Where("username = ? AND password=?", user.Username, user.Password).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "登录失败",
		})
		return
	}
	TokenString, _ := JWT.GenToken(user.Username)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": gin.H{
			"token": TokenString,
			"user":  user,
		},
	})

}
func Edituser(c *gin.Context) {
	id := c.Param("id")
	user := new(models.User)
	err := dao.DB.First(&user, id).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "该用户不存在",
		})
		return
	}
	c.BindJSON(&user)
	dao.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}
func Createpost(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)
	err := dao.DB.Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "创建错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": post,
	})

}
func Getposts(c *gin.Context) {

	var posts []models.Post
	Title := c.Query("Title")
	Page := c.Query("Page")
	PageNum, _ := strconv.Atoi(Page)
	Pagesize := c.Query("PageSize")
	PagesizeNum, _ := strconv.Atoi(Pagesize)

	if Title == "" {
		dao.DB.Offset((PageNum - 1) * PagesizeNum).Limit(PagesizeNum).Find(&posts)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": posts,
		})
		return
	} else {
		err := dao.DB.Where("Title=?", Title).Find(&posts).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "无数据",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": posts,
		})
	}

}
func Editpost(c *gin.Context) {
	id := c.Param("id")
	post := new(models.Post)
	dao.DB.First(&post, id)
	c.ShouldBind(&post)
	err := dao.DB.Save(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": post,
	})

}
func Deletepost(c *gin.Context) {
	var post models.Post
	c.ShouldBind(&post)
	err := dao.DB.Delete(&models.Post{}, post.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})

}
func Getpostnumber(c *gin.Context) {
	Title := c.Query("Title")
	var posts []models.Post
	if Title == "" {
		dao.DB.Find(&posts)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": posts,
		})
		return
	} else {
		err := dao.DB.Where("Title=?", Title).Find(&posts).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": posts,
		})
	}
}
