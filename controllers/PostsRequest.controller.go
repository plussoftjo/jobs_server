// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StorePostsRequest ..
func StorePostsRequest(c *gin.Context) {
	var data models.PostsRequest

	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var query models.PostsRequest
	config.DB.Where("id = ?", data.ID).First(&query)

	c.JSON(200, query)
}

// ShowPostsRequest ..
func ShowPostsRequest(c *gin.Context) {
	ID := c.Param("id")
	var data models.PostsRequest

	err := config.DB.Where("id = ?", ID).First(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// ShowNextPostsRequest ..
func ShowNextPostsRequest(c *gin.Context) {
	ID := c.Param("id")
	var data models.PostsRequest

	err := config.DB.Where("id > ?", ID).First(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// ShowPreviousPostsRequest ..
func ShowPreviousPostsRequest(c *gin.Context) {
	ID := c.Param("id")
	var data models.PostsRequest

	err := config.DB.Where("id < ?", ID).Last(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// ShowLastPostsRequest ..
func ShowLastPostsRequest(c *gin.Context) {
	var data models.PostsRequest

	err := config.DB.Last(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// ShowFirstPostsRequest ..
func ShowFirstPostsRequest(c *gin.Context) {
	var data models.PostsRequest

	err := config.DB.First(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	c.JSON(200, data)
}

// UpdatePostsRequest ..
func UpdatePostsRequest(c *gin.Context) {
	var data models.PostsRequest
	c.ShouldBindJSON(&data)

	err := config.DB.Save(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// RemovePostsRequest ..
func RemovePostsRequest(c *gin.Context) {
	ID := c.Param("id")

	err := config.DB.Delete(&models.PostsRequest{}, ID).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// IndexPostsRequest ..
func IndexPostsRequest(c *gin.Context) {
	var data []models.PostsRequest
	config.DB.Find(&data)

	c.JSON(200, data)
}
