// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreYoutubePosts ..
func StoreYoutubePosts(c *gin.Context) {
	var data models.YoutubePosts

	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var query models.YoutubePosts
	config.DB.Where("id = ?", data.ID).First(&query)

	c.JSON(200, query)
}

// ShowYoutubePosts ..
func ShowYoutubePosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.YoutubePosts

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

// ShowNextYoutubePosts ..
func ShowNextYoutubePosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.YoutubePosts

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

// ShowPreviousYoutubePosts ..
func ShowPreviousYoutubePosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.YoutubePosts

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

// ShowLastYoutubePosts ..
func ShowLastYoutubePosts(c *gin.Context) {
	var data models.YoutubePosts

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

// ShowFirstYoutubePosts ..
func ShowFirstYoutubePosts(c *gin.Context) {
	var data models.YoutubePosts

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

// UpdateYoutubePosts ..
func UpdateYoutubePosts(c *gin.Context) {
	var data models.YoutubePosts
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

// RemoveYoutubePosts ..
func RemoveYoutubePosts(c *gin.Context) {
	ID := c.Param("id")

	err := config.DB.Delete(&models.YoutubePosts{}, ID).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// IndexYoutubePosts ..
func IndexYoutubePosts(c *gin.Context) {
	var data []models.YoutubePosts
	config.DB.Order("id desc").Find(&data)

	c.JSON(200, data)
}
