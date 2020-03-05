package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}

		//博客创建
		v1.POST("blogs", api.CreateBlog)
		//博客详情
		v1.GET("blog/:id", api.ShowBlog)
		//博客列表
		v1.GET("blogs", api.ListBlog)
		//博客更新
		v1.PUT("blog/:id", api.UpdateBlog)
		//博客删除
		v1.DELETE("blog/:id", api.DeleteBlog)
	}
	return r
}
