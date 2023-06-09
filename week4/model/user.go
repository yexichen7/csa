package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	question string `json:"question"`
	answer   string `json:"answer"`
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
