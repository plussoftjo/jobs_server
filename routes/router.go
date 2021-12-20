// Package routes (Setup Routes Group)
package routes

import (
	"server/config"
	"server/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup >>>
func Setup() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "authorization", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))
	// gin.SetMode(gin.ReleaseMode)
	r.Use(static.Serve("/public", static.LocalFile(config.ServerInfo.PublicPath+"public", true)))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	// -------- Auth Groups ----------//

	// ~~~ Auth Group ~~~ //
	auth := r.Group("/auth")
	auth.POST("/login", controllers.LoginController)
	auth.POST("/register", controllers.RegisterController)
	auth.POST("/app/register", controllers.AppRegisterController)
	auth.POST("/app/login", controllers.AppLoginController)
	auth.GET("/app/auth", controllers.AuthAppUser)
	auth.POST("/app/changePassword", controllers.ChangePassword)
	auth.GET("/auth", controllers.Auth)
	auth.GET("/users/index", controllers.UsersListIndex)
	auth.GET("/users/delete/:id", controllers.DeleteUser)
	auth.POST("/update", controllers.UpdateUser)
	auth.POST("/app/update", controllers.AppUpdateUser)
	auth.POST("/checkHasPhone", controllers.CheckIfHasPhone)
	auth.POST("/resetPassword", controllers.ResetPassword)

	// --------- Basics ------- //
	basics := r.Group("/basics")

	// UploadImage => For All
	basics.POST("/upload_image/:imageType", controllers.UpdateImage)

	// --------- User Controller ----------------- //
	user := r.Group("/users")
	// ~~~ User Roles ~~~ //
	user.POST("/roles/store", controllers.StoreUserRoles)
	user.POST("/roles/update", controllers.UpdateUserRole)
	user.GET("/roles/index", controllers.IndexUserRoles)
	user.GET("/roles/delete/:id", controllers.DeleteUserRole)
	// --------------- Employ Controller ----------- //
	user.POST("/employee/store", controllers.StoreEmployee)
	user.GET("/employee/index", controllers.IndexEmployee)
	user.GET("/employee/delete/:id", controllers.DeleteEmployee)
	user.POST("/employee/update", controllers.UpdateEmployee)

	// Dashboard
	dashboard := r.Group("/dashboard")
	dashboard.GET("/indexAllClients", controllers.IndexAllClients)
	dashboard.GET("/showUser/:id", controllers.ShowUser)

	// --------- Posts ------- //
	Posts := r.Group("/Posts")
	Posts.POST("/store", controllers.StorePosts)
	Posts.POST("/update", controllers.UpdatePosts)
	Posts.GET("/show/:id", controllers.ShowPosts)
	Posts.GET("/remove/:id", controllers.RemovePosts)
	Posts.GET("/showPrevious/:id", controllers.ShowPreviousPosts)
	Posts.GET("/showNext/:id", controllers.ShowNextPosts)
	Posts.GET("/showFirst", controllers.ShowFirstPosts)
	Posts.GET("/showLast", controllers.ShowLastPosts)
	Posts.GET("/index", controllers.IndexPosts)
	Posts.GET("/pagination/:page", controllers.IndexPagination)
	Posts.POST("/filter/:page", controllers.FilterResults)
	Posts.GET("/index/rand", controllers.RandIndexPosts)

	// --------- YoutubePosts ------- //
	YoutubePosts := r.Group("/YoutubePosts")
	YoutubePosts.POST("/store", controllers.StoreYoutubePosts)
	YoutubePosts.POST("/update", controllers.UpdateYoutubePosts)
	YoutubePosts.GET("/show/:id", controllers.ShowYoutubePosts)
	YoutubePosts.GET("/remove/:id", controllers.RemoveYoutubePosts)
	YoutubePosts.GET("/showPrevious/:id", controllers.ShowPreviousYoutubePosts)
	YoutubePosts.GET("/showNext/:id", controllers.ShowNextYoutubePosts)
	YoutubePosts.GET("/showFirst", controllers.ShowFirstYoutubePosts)
	YoutubePosts.GET("/showLast", controllers.ShowLastYoutubePosts)
	YoutubePosts.GET("/index", controllers.IndexYoutubePosts)

	// --------- PostsRequest ------- //
	PostsRequest := r.Group("/PostsRequest")
	PostsRequest.POST("/store", controllers.StorePostsRequest)
	PostsRequest.POST("/update", controllers.UpdatePostsRequest)
	PostsRequest.GET("/show/:id", controllers.ShowPostsRequest)
	PostsRequest.GET("/remove/:id", controllers.RemovePostsRequest)
	PostsRequest.GET("/showPrevious/:id", controllers.ShowPreviousPostsRequest)
	PostsRequest.GET("/showNext/:id", controllers.ShowNextPostsRequest)
	PostsRequest.GET("/showFirst", controllers.ShowFirstPostsRequest)
	PostsRequest.GET("/showLast", controllers.ShowLastPostsRequest)
	PostsRequest.GET("/index", controllers.IndexPostsRequest)

	// --------- Tags ------- //
	Tags := r.Group("/Tags")
	Tags.POST("/store", controllers.StoreTags)
	Tags.POST("/update", controllers.UpdateTags)
	Tags.GET("/show/:id", controllers.ShowTags)
	Tags.GET("/remove/:id", controllers.RemoveTags)
	Tags.GET("/showPrevious/:id", controllers.ShowPreviousTags)
	Tags.GET("/showNext/:id", controllers.ShowNextTags)
	Tags.GET("/showFirst", controllers.ShowFirstTags)
	Tags.GET("/showLast", controllers.ShowLastTags)
	Tags.GET("/index", controllers.IndexTags)

	notifications := r.Group("/notifications")
	notifications.POST("/store", controllers.StoreNotification)
	notifications.POST("/storeNotification", controllers.StoreNotificationToken)

	r.Run(":8082")
}
