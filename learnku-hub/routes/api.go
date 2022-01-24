// Package routes 注册路由
package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API（根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			// 注册
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			// 使用手机注册
			authGroup.POST("/signup/using-phone",
				middlewares.GuestJWT(),
				middlewares.LimitPerRoute("60-H"),
				suc.SignupUsingPhone)
			// 使用邮箱注册
			authGroup.POST("/signup/using-email",
				middlewares.GuestJWT(),
				middlewares.LimitPerRoute("60-H"),
				suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)

			// 登录
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 使用手机, Email, 用户名 和密码进行登录
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// refresh token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.AuthJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.AuthJWT(), pwc.ResetByEmail)
		}
	}

	uc := new(controllers.UsersController)
	// 获取当前用户
	v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
	usersGroup := v1.Group("/users")
	{
		usersGroup.GET("", uc.Index)
	}

	cgc := new(controllers.CategoriesController)
	cgcGroup := v1.Group("/categories")
	{
		cgcGroup.GET("", cgc.Index)
		cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
		cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
		cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
	}
}
