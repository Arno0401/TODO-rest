package server

import (
	config "arno/configs"
	handler "arno/internal/handlers"
	"arno/middleware"
	"github.com/gin-gonic/gin"
	"net"
)

type Server struct {
	g *gin.Engine
}

func New(handler *handler.Handler) *Server {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	r.POST("/sign_up", handler.SignUpHandler)
	r.POST("/sign_in", handler.SignInHandler)
	r.GET("/profile", handler.Profile)
	r.POST("/change_password", handler.ChangePassword)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/todos", handler.CreateTodos)
		authRoutes.GET("/todos", handler.GetTodos)
		authRoutes.PUT("/todos/:id", handler.UpdateTodos)
		authRoutes.DELETE("/todos/:id", handler.DeleteTodo)
	}

	return &Server{
		g: r,
	}
}

func (s *Server) Run() error {
	addr := net.JoinHostPort(config.DBConfig.Server.Host, config.DBConfig.Server.Port)
	return s.g.Run(addr)
}
