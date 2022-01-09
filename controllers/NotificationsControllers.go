package controllers

import (
	"fmt"
	"server/config"
	"server/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

type storeNotificationTypes struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ForPost  bool   `json:"forPost"`
	PostID   uint   `json:"postID"`
	ForUsers bool   `json:"forUsers"`
	UsersID  []int  `json:"usersID"`
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

// StoreNotifications ..
func StoreNotification(c *gin.Context) {
	// Get the data
	var data storeNotificationTypes
	c.ShouldBindJSON(&data)
	// End getting the data

	var usersToken []string
	if data.ForUsers {
		config.DB.Model(&models.NotificationTokens{}).Where("id IN ("+arrayToString(data.UsersID, ",")+")").Pluck("token", &usersToken)
	} else {
		config.DB.Model(&models.NotificationTokens{}).Pluck("token", &usersToken)
	}

	var toPushTokens []expo.ExponentPushToken
	client := expo.NewPushClient(nil)
	dataType := "main"
	postID := 0
	if data.ForPost {
		dataType = "post"
		postID = int(data.PostID)
	}

	for _, tok := range usersToken {
		pushToken, err := expo.NewExponentPushToken(tok)
		if err != nil {
			fmt.Println("Error notification id")
			return
		}

		// Publish message
		response, err := client.Publish(
			&expo.PushMessage{
				To:       []expo.ExponentPushToken{pushToken},
				Body:     data.Body,
				Data:     map[string]string{"type": dataType, "postID": strconv.Itoa(postID)},
				Sound:    "default",
				Title:    data.Title,
				Priority: expo.DefaultPriority,
			},
		)

		// Check errors
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Validate responses
		if response.ValidateResponse() != nil {
			c.JSON(500, gin.H{
				"error": "There are error sending",
			})
			return
		}

	}

	c.JSON(200, toPushTokens)
}

func StoreNotificationToken(c *gin.Context) {
	var data models.NotificationTokens
	c.ShouldBindJSON(&data)

	config.DB.Create(&data)

	c.JSON(200, gin.H{"message": "success"})
}
