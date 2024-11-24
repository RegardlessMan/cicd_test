/**
 * @Author QG
 * @Date  2024/11/24 16:08
 * @description
**/

package controllers

import (
	"cicd_test/global"
	"cicd_test/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateExchangeRate
//
//	@Description: 创建exchange_rate
//	@param c
func CreateExchangeRate(c *gin.Context) {
	var exchangeRate models.ExchangeRate

	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Migrate Error"})
		return
	}

	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create Exchange Rate Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Exchange Rate Created Successfully"})
}
