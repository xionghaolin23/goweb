package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
gin快速入门
https://zhuanlan.zhihu.com/p/375688504
*/
func main() {

	//创建一个服务
	server := gin.Default()
	//server.Get("/", ...)声明了一个路由
	server.GET("/hello", func(context *gin.Context) {
		context.String(200, "Hello, world")
	})

	//1、获取query参数

	//匹配 http://localhost:8080/user?name=zhangsan&age=20
	server.GET("/user", func(context *gin.Context) {
		name, _ := context.GetQuery("name")
		age, _ := context.GetQuery("age")
		context.String(http.StatusOK, "%s is a %s", name, age)
	})

	//2、url传参
	//匹配 http://localhost:8080/user/zhangsan
	server.GET("user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s", name)
	})

	//3、获取POST参数（body参数）
	//{"password":"1234","username":"geektutu"}
	server.POST("user", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	//服务端口
	server.Run()

}
