/**
 * @Author QG
 * @Date  2024/11/25 23:17
 * @description
**/

package controllers

import (
	"cicd_test/global"
	"cicd_test/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ok := global.Db.AutoMigrate(&article); ok != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database Migrate Error"})
	}

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Create Article Error"})
		return
	}

	ctx.JSON(http.StatusCreated, article)
}

func GetArticles(ctx *gin.Context) {

	var articles []models.Article

	if err := global.Db.Find(&articles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, articles)

}

func GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")

	var article models.Article
	if err := global.Db.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)
}

func LikeArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	first := global.Db.First(&models.Article{}, id)
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	global.RedisDb.Incr("article:" + id + ":likes")
	ctx.JSON(http.StatusOK, gin.H{"message": "Article liked successfully"})
}

func UnlikeArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	first := global.Db.First(&models.Article{}, id)
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	global.RedisDb.Decr("article:" + id + ":likes")
	ctx.JSON(http.StatusOK, gin.H{"message": "Article unliked successfully"})
}
