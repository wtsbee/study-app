package router

import (
	"net/http"
	"os"
	"study-app-api/controller"
	"study-app-api/logging"

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
	e.POST("/signup", logging.LoggingMiddleware(uc.SignUp))
	e.POST("/login", logging.LoggingMiddleware(uc.LogIn))
	e.POST("/logout", logging.LoggingMiddleware(uc.LogOut))
	e.GET("/csrf", logging.LoggingMiddleware(uc.CsrfToken))
	t := e.Group("/task")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	t.POST("/upload", logging.LoggingMiddleware(tc.UploadImageHandler))
	t.Static("/images", "images")
	t.GET("/:id", logging.LoggingMiddleware(tc.GetTask))
	t.POST("", logging.LoggingMiddleware(tc.CreateTask))
	t.POST("/:id", logging.LoggingMiddleware(tc.UpdateTask))
	t.DELETE("/:id", logging.LoggingMiddleware(tc.DeleteTask))
	t.GET("/:id/detail", logging.LoggingMiddleware(tdc.GetTaskDetail))
	t.POST("/:id/detail", logging.LoggingMiddleware(tdc.CreateTaskDetail))
	t.POST("/:id/update_task_detail", logging.LoggingMiddleware(tdc.UpdateTaskDetail))
	t.GET("/:id/ws", logging.LoggingMiddleware(tdc.WebSocketHandlerForTaskDetail))
	ts := e.Group("/tasks")
	ts.GET("/ws", logging.LoggingMiddleware(tc.WebSocketHandler))
	ts.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	ts.GET("", logging.LoggingMiddleware(tc.GetOwnAllTasks))
	ts.POST("", logging.LoggingMiddleware(tc.UpdateOwnAllTasks))
	ts.DELETE("/:id", logging.LoggingMiddleware(tc.DeleteTaskList))
	return e
}
