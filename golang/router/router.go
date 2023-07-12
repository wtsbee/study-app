package router

import (
	"net/http"
	"os"
	"study-app-api/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, tdc controller.ITaskDetailController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		// CookieSameSite: http.SameSiteDefaultMode,
		//CookieMaxAge:   60,
	}))
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/task")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	t.POST("/upload", tc.UploadImageHandler)
	t.Static("/images", "images")
	t.GET("/:id", tc.GetTask)
	t.POST("", tc.CreateTask)
	t.POST("/:id", tc.UpdateTask)
	t.GET("/:id/detail", tdc.GetTaskDetail)
	t.POST("/:id/detail", tdc.CreateTaskDetail)
	t.POST("/:id/update_task_detail", tdc.UpdateTaskDetail)
	t.GET("/:id/ws", tdc.WebSocketHandlerForTaskDetail)
	ts := e.Group("/tasks")
	ts.GET("/ws", tc.WebSocketHandler)
	ts.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	ts.GET("", tc.GetOwnAllTasks)
	ts.POST("", tc.UpdateOwnAllTasks)
	ts.DELETE("/:id", tc.DeleteTaskList)
	return e
}
