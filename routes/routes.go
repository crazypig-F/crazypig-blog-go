package routes

import (
	"BlogProject/controller"
	"BlogProject/controller/front"
	"BlogProject/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	//store := cookie.NewStore([]byte("something-very-secret"))
	//r.Use(sessions.Sessions("my_session", store))
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//上传操作
		v1.POST("upload", controller.UpLoad)
		//前台用户
		v1.GET("front/blogs", front.ListFrontBlogController)
		v1.GET("front/tags", front.ListFrontTagController)
		v1.GET("front/types", front.ListFrontTypeController)
		v1.GET("front/blog/:id", front.ShowFrontBlogController)
		// 登陆注册操作
		v1.POST("user/register", controller.UserRegisterController)
		v1.POST("user/login", controller.UserLoginController)
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//博客管理
			authed.GET("blogs", controller.ListBlogController)
			authed.POST("blog", controller.CreateBlogController)
			authed.GET("blog/:id", controller.ShowBlogController)
			authed.DELETE("blog/:id", controller.DeleteBlogController)
			authed.PUT("blog/:id", controller.UpdateBlogController)
			authed.POST("blog/search", controller.SearchBlogController)
			//类型管理
			authed.GET("types", controller.ListTypeController)
			authed.POST("type", controller.CreateTypeController)
			authed.GET("type/:id", controller.ShowTypeController)
			authed.DELETE("type/:id", controller.DeleteTypeController)
			authed.PUT("type/:id", controller.UpdateTypeController)
			authed.POST("type/search", controller.SearchTypeController)
			// 标签管理
			authed.GET("tags", controller.ListTagController)
			authed.POST("tag", controller.CreateTagController)
			authed.GET("tag/:id", controller.ShowTagController)
			authed.DELETE("tag/:id", controller.DeleteTagController)
			authed.PUT("tag/:id", controller.UpdateTagController)
			authed.POST("tag/search", controller.SearchTagController)
		}
	}
	return r
}
