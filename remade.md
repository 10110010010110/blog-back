markdown
# 使用 Gin 框架实现简易博客系统后端搭建步骤

## 一、创建 Golang 项目
1. 在开发工具中点击新建项目
2. 配置 Go 环境代理（加速依赖下载）：
   ```bash
   go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
二、下载 Gin 框架
执行以下命令安装 Gin 框架：
bash
go get -u github.com/gin-gonic/gin
三、下载 GORM（ORM 框架）
执行以下命令安装 GORM：
bash
go get -u gorm.io/gorm
四、连接 MySQL 数据库
(1) 下载 MySQL 驱动包
bash
go get -u gorm.io/driver/mysql
(2) 数据库连接代码示例
go
运行
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func main() {
  // 数据库连接字符串（格式：用户名:密码@tcp(地址:端口)/数据库名?参数）
  dsn := "root:root@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
  
  // 连接数据库
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("连接数据库失败: " + err.Error())
  }
  
  // 成功连接后，db 变量可用于后续数据库操作
}
五、下载 JWT（用于身份认证）
执行以下命令安装 JWT 库：
bash
go get github.com/golang-jwt/jwt/v4
 

  
  
  

  