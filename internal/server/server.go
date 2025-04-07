package server

import (
	"ToGoList/internal/config"
	"ToGoList/internal/routes"
	"ToGoList/pkg/database"
	"ToGoList/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *database.DB
}

func NewServer(cfg *config.Config) *Server {
	db := database.InitDB(cfg.DbConnectionString)

	r := gin.Default()
	r.Use(middleware.DbMiddleware(db))

	routes.SetupRoutes(r)

	return &Server{
		router: r,
		cfg:    cfg,
		db:     db,
	}
}

func (s *Server) Run() error {
	return s.router.Run(s.cfg.ServerHost)
}
