      
## 一、创建 Golang 项目

1. 在开发工具中点击新建项目

2. 配置 Go 环境代理（加速依赖下载）：

   ```bash
   go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
   ```

## 二、下载gin框架

```bash
go get -u github.com/gin-gonic/gin
```

## 三、下载grom

```bash
go get -u gorm.io/gorm
```

## 四、连接数据库（以mysql为例）

### (1) 下载mysql驱动包

```bash
go get -u gorm.io/driver/mysql
```

### (2) 连接数据库

```go
dsn := "root:root@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

## 五、下载jwt

```bash
go get github.com/golang-jwt/jwt/v4
```
        
   