// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"

	"github.com/gin-gonic/gin"
)

// StoreTags ..
func StoreTags(c *gin.Context) {
	var data models.Tags

	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var query models.Tags
	config.DB.Where("id = ?", data.ID).First(&query)

	c.JSON(200, query)
}

// ShowTags ..
func ShowTags(c *gin.Context) {
	ID := c.Param("id")
	var data models.Tags

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

// ShowNextTags ..
func ShowNextTags(c *gin.Context) {
	ID := c.Param("id")
	var data models.Tags

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

// ShowPreviousTags ..
func ShowPreviousTags(c *gin.Context) {
	ID := c.Param("id")
	var data models.Tags

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

// ShowLastTags ..
func ShowLastTags(c *gin.Context) {
	var data models.Tags

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

// ShowFirstTags ..
func ShowFirstTags(c *gin.Context) {
	var data models.Tags

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

// UpdateTags ..
func UpdateTags(c *gin.Context) {
	var data models.Tags
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

// RemoveTags ..
func RemoveTags(c *gin.Context) {
	ID := c.Param("id")

	err := config.DB.Delete(&models.Tags{}, ID).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// IndexTags ..
func IndexTags(c *gin.Context) {
	var data []models.Tags
	config.DB.Find(&data)

	c.JSON(200, data)
}
