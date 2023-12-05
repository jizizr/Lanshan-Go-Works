package boot

import (
	"ezgin/controller"
	"ezgin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.Cors)
	public := r.Group("")

	{
		public.POST("/register", controller.Register)
		public.POST("/reset", controller.ResetPwd)
		public.POST("/login", controller.Login)
		public.GET("/comment/:uid", controller.GetComment)
	}

	private := r.Group("")
	private.Use(middleware.JWTAuth)
	{
		private.POST("/comment", controller.PostComment)
	}

	r.Run(":8080")
}
