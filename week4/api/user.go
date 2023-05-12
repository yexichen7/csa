package api

import (
	"csa/week4/api/middleware"
	"csa/week4/dao"
	"csa/week4/model"
	"csa/week4/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "verification failed")
		return
	}
	// 传入用户名和密码,并获取密保问题
	username := c.PostForm("username")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	// 验证用户名是否重复
	flag := dao.SelectUser(username)
	fmt.Println(flag)
	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user already exists")
		return
	}

	dao.AddUser(username, password, question, answer)
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

// 仅有登录部分有改动
func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != password {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}

	// 正确则登录成功
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "Yxh",                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, tokenString)
}
func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func change(c *gin.Context) {

	// 传入用户名和密码
	username := c.PostForm("username")
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")
	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	// 查找正确的密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 若不正确则传出错误
	if selectPassword != oldPassword {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong password")
		return
	}
	ok := dao.Change(username, oldPassword, newPassword)
	if !ok {
		utils.RespFail(c, "change password fail")
	}
	utils.RespSuccess(c, "change password success")
}

func findPassword(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 传入用户名和密码
	username := c.PostForm("username")
	answer := c.PostForm("secretQuestion")

	// 验证用户名是否存在
	flag := dao.SelectUser(username)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	// 查找正确的答案
	selectAnswer := dao.SelectAnswerFromUsername(username)
	// 若不正确则传出错误
	if selectAnswer != answer {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "wrong answer")
		return
	}

	// 正确则找回密码
	selectPassword := dao.SelectPasswordFromUsername(username)
	// 创建一个我们自己的声明
	utils.RespSuccess(c, selectPassword)
}
