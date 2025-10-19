golang gin框架实现简易博客系统后端部分
一.创建golang项目
  在文件处点击新建，找到项目，然后用阿里云代理：
  go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
二.下载gin框架
  go get -u github.com/gin-gonic/gin
三.下载grom
  go get -u gorm.io/gorm
四.连接数据库（以mysql为例）
  (1)下载mysql驱动包
  go get -u gorm.io/driver/mysql
  (2)连接数据库
  dsn := "root:root@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
  db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
五.下载jwt
  go get github.com/golang-jwt/jwt/v4  
 

  
  
  

  