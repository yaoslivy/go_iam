package apiserver

import (
	"go_iam/internal/apiserver/controller/v1/secret"
	"go_iam/internal/apiserver/controller/v1/user"
	"go_iam/internal/apiserver/store/mysql"
	"go_iam/internal/pkg/code"
	"go_iam/internal/pkg/middleware"
	"go_iam/internal/pkg/middleware/auth"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"

	// custom gin validators.
	_ "go_iam/pkg/validator"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {
}

func installController(g *gin.Engine) *gin.Engine {
	// Middlewares.
	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler) // Basic authentication
	// Refresh time can be longer than token timeout
	g.POST("/refresh", jwtStrategy.RefreshHandler) // Bearer authentication
	g.POST("/logout", jwtStrategy.LogoutHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found."), nil)
	})

	//v1 handlers, requiring authentication
	storeIns, _ := mysql.GetMySQLFactoryOr(nil) // Get mysql store
	v1 := g.Group("/v1")
	{
		v1.Use(auto.AuthFunc()) // auto authentication strategy.

		//user RESTful resource
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)

			// Add admin validation for this group.
			// userv1.Use(auto.AuthFunc(), middleware.Validation())
			userv1.Use(middleware.Validation())

			userv1.POST("", userController.Create)
			userv1.GET(":name", userController.Get)    //(admin api)
			userv1.PUT(":name", userController.Update) //update the data, except for the password
			userv1.PUT(":name/change-password", userController.ChangePassword)
			userv1.DELETE(":name", userController.Delete) //(admin api)

			userv1.GET("", userController.List)                //BindQuery
			userv1.DELETE("", userController.DeleteCollection) //(admin api)
		}

		//secret RESTful resource
		secretv1 := v1.Group("/secrets")
		{
			//secret controller
			secretController := secret.NewSecretController(storeIns)

			secretv1.GET("", secretController.List)
			secretv1.POST("", secretController.Create)
			secretv1.GET(":name", secretController.Get)
		}
	}

	return g
}
