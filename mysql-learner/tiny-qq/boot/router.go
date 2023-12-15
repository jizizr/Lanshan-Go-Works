package boot

import (
	"github.com/gin-gonic/gin"
	"tiny-qq/controller"
	"tiny-qq/middleware"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.Cors)
	public := r.Group("")

	{
		public.POST("/register", controller.Register)
		public.POST("/login", controller.Login)
	}

	private := r.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.POST("/friend", controller.AddFriend)
		private.DELETE("/friend", controller.DeleteFriend)
		private.GET("/friend", controller.QueryFriendsList)
		private.POST("/group", controller.CreateGroup)
		private.DELETE("/group", controller.DeleteGroup)
		private.GET("/group", controller.QueryGroupsList)
		private.POST("/grouper", controller.AddGroupUser)
		private.DELETE("/grouper", controller.DeleteGroupUser)
		private.GET("/search", controller.SearchFriend)
	}

	r.Run(":8080")
}
