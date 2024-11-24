package controllers

import (
	"cicd_test/global"
	"cicd_test/models"
	"cicd_test/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid JSON",
		})
		return
	}

	//密码加密
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Hash Password Error",
		})
		return
	}

	user.Password = hashPassword

	if err = global.Db.AutoMigrate(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Database Migrate Error",
		})
		return
	}
	if err = global.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Create User Error",
		})
		return
	}

	//生成token
	jwtToken, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Generate JWT Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Invalid JSON",
		})
		return
	}

	var user models.User
	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "User Not Exist",
		})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Password Not Correct",
		})
		return
	}

	jwtToken, _ := utils.GenerateJWT(user.Username)

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
