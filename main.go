package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//传统的Web服务写法
	//http.HandleFunc("/hello", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//    fmt.Printf("http server faile,err:%v\n", err)
	//    return
	//}
	//fmt.Println("项目启动成功")

	//利用Gin框架的web写法，来源于gin官网
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
	panic(r.Run())
}
