package server

import (
	"myblog/api"
	"myblog/middleware"
	"os"

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

	dispatchRouterLogin(r)
	dispatchRouterBlog(r)
	dispatchRouterTag(r)

	return r
}

func dispatchRouterLogin(router *gin.Engine) {
	groups := router.Group("/api/v1")
	{
		groups.POST("ping", api.Ping)

		// 用户注册
		groups.POST("user/register", api.UserRegister)
		// 用户登录
		groups.POST("user/login", api.UserLogin)
		// 需要登录保护的
		auth := groups.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
}
func dispatchRouterBlog(router *gin.Engine) {
	groups := router.Group("/api/v1")
	{
		//博客创建
		groups.POST("blogs", api.CreateBlog)
		//博客详情
		groups.GET("blog/:id", api.ShowBlog)
		//博客列表
		groups.GET("blogs", api.ListBlog)
		//博客更新
		groups.PUT("blog/:id", api.UpdateBlog)
		//博客删除
		groups.DELETE("blog/:id", api.DeleteBlog)
	}
}

func dispatchRouterTag(router *gin.Engine) {
	groups := router.Group("/api/v1")
	{
		//添加标签
		groups.POST("tags", api.CreateTag)
		//获取标签
		groups.GET("tag/:id", api.ShowTag)
		//标签列表
		groups.GET("tags", api.ListTag)
		//修改标签
		groups.PUT("tag/:id", api.UpdateTag)
		//删除标签
		groups.DELETE("tag/:id", api.DeleteTag)
	}
}
