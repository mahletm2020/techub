package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func Signup(c *gin.Context){	
//getting email fromreq body

var body struct{
	UserName string
	email    string
	paswword string
	role     string
}
//to populate the variable with the data comin in  
c.Bind(&body)!= nil{
	c.JSON(http.StatusBadRequest, gin.H{
		"error":"failed to raad body"
		,
	})
	return
}
//we have he body with its user password exsetera exetra so lets hash he password 

hash,err := bcrypt.GenerateFrompassword([]byte(body.password),10)

if err != nil{
	c.JSON(http.StatusBadRequest,gin.H{
		"error":"failed to hash password",
	})
	return
}
//creatin user
user:=

}
