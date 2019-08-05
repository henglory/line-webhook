package server

import (
	"context"
	golog "log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/henglory/line-webhook/conf"
	"github.com/henglory/line-webhook/handler"
)

type errorResponse struct {
	ResponseCode int64  `json:"responseCode"`
	Reason       string `json:"reason"`
	RawRequest   string `json:"rawRequest"`
}

type Server struct {
	srv *http.Server
}

func NewServer() *Server {
	server := &Server{}
	return server
}

func (server *Server) Start() {
	go server.ginStart()
}

func (server *Server) Close() {
	server.srv.Close()
}

type readiness struct {
	Success bool `json:"success"`
}

func (server Server) ginStart() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(500, readiness{
				Success: false,
			})
			return
		}
		handler.Webhook(b)
		c.JSON(200, readiness{
			Success: true,
		})
	})

	r.POST("/", func(c *gin.Context) {
		b, err := c.GetRawData()
		if err != nil {
			c.JSON(500, readiness{
				Success: false,
			})
			return
		}
		handler.Webhook(b)
		c.JSON(200, readiness{
			Success: true,
		})
	})

	server.srv = &http.Server{
		Addr:    conf.ServicePort,
		Handler: r,
	}

	if err := server.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		golog.Fatalf("listen: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.srv.Shutdown(ctx); err != nil {
		golog.Fatal("Server Shutdown:", err)
	}
	defer cancel()
}
