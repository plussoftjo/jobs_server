// Package controllers ..
package controllers

import (
	"server/config"
	"server/models"
	"server/vendors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StorePosts ..
func StorePosts(c *gin.Context) {
	var data models.Posts

	c.ShouldBindJSON(&data)

	err := config.DB.Create(&data).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var query models.Posts
	config.DB.Where("id = ?", data.ID).First(&query)

	c.JSON(200, query)
}

// ShowPosts ..
func ShowPosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.Posts

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

// ShowNextPosts ..
func ShowNextPosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.Posts

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

// ShowPreviousPosts ..
func ShowPreviousPosts(c *gin.Context) {
	ID := c.Param("id")
	var data models.Posts

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

// ShowLastPosts ..
func ShowLastPosts(c *gin.Context) {
	var data models.Posts

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

// ShowFirstPosts ..
func ShowFirstPosts(c *gin.Context) {
	var data models.Posts

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

// UpdatePosts ..
func UpdatePosts(c *gin.Context) {
	var data models.Posts
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

// RemovePosts ..
func RemovePosts(c *gin.Context) {
	ID := c.Param("id")

	err := config.DB.Delete(&models.Posts{}, ID).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

// IndexPosts ..
func IndexPosts(c *gin.Context) {
	var data []models.Posts
	config.DB.Order("id desc").Find(&data)

	c.JSON(200, data)
}

// IndexPagination ..
func IndexPagination(c *gin.Context) {
	var data []models.Posts
	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	config.DB.Scopes(vendors.Paginate(page)).Order("id desc").Find(&data)

}

type FilterTypes struct {
	Search      string `json:"search"`
	Nationality string `json"nationality"`
	Tags        string `json:"tags"`
}

// FilterResults ..
func FilterResults(c *gin.Context) {

	var filter FilterTypes
	c.ShouldBindJSON(&filter)

	pageString := c.Param("page")
	page, err := strconv.Atoi(pageString)
	if err != nil {
		c.JSON(500, gin.H{
			"err":  err.Error(),
			"code": 101,
		})
		return
	}

	var data []models.Posts
	config.DB.
		Where("title LIKE ?", "%"+filter.Search+"%").
		Where("nationality = ?", filter.Nationality).
		Scopes(vendors.Paginate(page)).Find(&data)

	c.JSON(200, data)
}

// RandIndexPosts ..
func RandIndexPosts(c *gin.Context) {
	var data []models.Posts

	config.DB.Order("RAND()").Limit(8).Find(&data)

	c.JSON(200, data)

}
