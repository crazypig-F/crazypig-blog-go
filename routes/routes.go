package routes

import (
	"BlogProject/controller/qiniu"
	"BlogProject/controller/user"
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
		v1.POST("upload", qiniu.UpLoad)
		// 用户操作
		v1.POST("user/register", user.RegisterController)
		v1.POST("user/login", user.LoginController)
		authed := v1.Group("/") //需要登陆保护
		authed.Use(middleware.JWT())
		{
			//任务操作
			//authed.GET("tasks", api.ListTasks)
			//authed.POST("task", api.CreateTask)
			//authed.GET("task/:id", api.ShowTask)
			//authed.DELETE("task/:id", api.DeleteTask)
			//authed.PUT("task/:id", api.UpdateTask)
			//authed.POST("search", api.SearchTasks)
		}
	}
	return r
}
