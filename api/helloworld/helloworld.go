package helloworld

import (
	fmt "fmt"

	"github.com/gin-gonic/gin"
)

// @Success 200 {string} string
// @Router /demo/v1/hello [get]
func Hello(c *gin.Context) {
	a := 1
	fmt.Printf("%d", a)
	c.JSON(200, "Hello - v1")
}

// @Success 200 {string} string
// @Router /demo/v1/hi [post]
func Hi(c *gin.Context) {
	var cMap map[string]string

	cMap = make(map[string]string)

	cMap["name"] = "sam.cheng"

	c.JSON(200, cMap)

	// c.JSON(200, "Hi - v")
}

// @Summary 註冊
// @Description 註冊
// @Tags User
// @Version 1.0.0
// @Produce json
// @Param name formData string true "用戶帳號"
// @Param password formData string true "用戶密碼"
// @Success 201 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/signup [post]
func Signup(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")
	var cMap map[string]string
	cMap = make(map[string]string)
	cMap["name"] = name
	cMap["password"] = pwd
	c.JSON(201, cMap)

	// var cMap map[string]string
	// cMap = make(map[string]string)
	// cMap["statusCode"] = "400"
	// cMap["error"] = "error"
	// cMap["message"] = "error message"
	// c.JSON(400, cMap)
}

// @Summary 登入
// @Description 登入
// @Tags User
// @Version 1.0.0
// @Produce json
// @Param name formData string true "用戶帳號"
// @Param password formData string true "用戶密碼"
// @Success 201 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func Login(c *gin.Context) {
	name := c.PostForm("name")
	pwd := c.PostForm("password")
	var cMap map[string]string
	cMap = make(map[string]string)
	cMap["name"] = name
	cMap["password"] = pwd
	c.JSON(201, cMap)

	// var cMap map[string]string
	// cMap = make(map[string]string)
	// cMap["statusCode"] = "400"
	// cMap["error"] = "error"
	// cMap["message"] = "error message"
	// c.JSON(400, cMap)
}
