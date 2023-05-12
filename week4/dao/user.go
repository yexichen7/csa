package dao

import _ "csa/week4/model"

// 假数据库，用 map 实现
// 懒得开个数据库，我偷
var database = map[string]string{
	"yxh": "123456",
	"wx":  "654321",
}
var questions = map[string]string{}

var answers = map[string]string{}

func AddUser(username, password, question, answer string) {
	//果然不用数据库好抽象啊
	database[username] = password
	questions[username] = question
	answers[question] = answer
}

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

func SelectAnswerFromUsername(username string) string {
	return answers[questions[username]]
}

func Change(username string, oldPassword string, newPassword string) bool {
	database[username] = newPassword
	if database[username] == oldPassword {
		return false
	}
	return true
}
