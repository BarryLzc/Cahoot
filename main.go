package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 此方法在该项目中未用，不用管
func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

type User struct {
	gorm.Model
	//'gorm:"type:varchar(20);not null"'
	Name     string
	Phone    string
	Password string
}

func main() {
	db := InitDB()
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
		//获取参数
		name := c.PostForm("name")
		phone := c.PostForm("phone")
		password := c.PostForm("password")
		//数据验证
		if len(phone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号格式不正确"})
			return
		}

		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
			return
		}
		if len(name) == 0 {
			name = RandomString(10)
			return
		}
		log.Print(name, phone, password)

		//判断手机号是否存在
		if isPhoneExist(db, phone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在，不能注册"})
			return
		}

		//创建新用户
		newUser := User{
			Name:     name,
			Phone:    phone,
			Password: password,
		}
		db.Create(&newUser)

		//返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080
	panic(r.Run())

}
func isPhoneExist(db *gorm.DB, phone string) bool {
	var user User
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 随机产生英文字符，可设定长度
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]

	}

	return string(result)
}

func InitDB() *gorm.DB {
	//前提是你要先在本机用Navicat创建一个名为go_db的数据库
	host := "localhost"
	port := "3306"
	database := "cahoot"
	username := "root"
	password := "LDnbV50104"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	//这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
	}
	//这个是gorm自动创建数据表的函数。它会自动在数据库中创建一个名为users的数据表
	_ = db.AutoMigrate(&User{})
	return db
}
