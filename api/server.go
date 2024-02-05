package api

import (
	"fmt"
	"net/http"
	"simple-dating-app-service/api/route"
	"simple-dating-app-service/config"
	db "simple-dating-app-service/db/sqlc"
	"simple-dating-app-service/domain"
	"simple-dating-app-service/token"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	domain.ConcreteServer
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(env config.Env, store db.Store, timeout time.Duration) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.Env.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)

	}

	server := &Server{
		ConcreteServer: domain.ConcreteServer{Store: store, Config: env, Timeout: timeout, TokenMaker: tokenMaker},
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	if server.Config.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	allowedUrls := []string{
		server.Config.BeliesimWebFrontendUrl,
		server.Config.XenditHostUrl,
		server.Config.UsimsaHostUrl,
	}

	// var allowedDomains []string

	// for _, url := range allowedUrls {
	// 	domain, err := util.GetUrlDomain(url)
	// 	if err != nil {
	// 		log.Fatalf("Error on util.GetUrlDomain : %v", err)
	// 	}
	// 	allowedDomains = append(allowedDomains, domain)
	// }

	config := cors.DefaultConfig()

	if server.Config.AppEnv == "prod" {
		config.AllowOrigins = allowedUrls
	} else {
		config.AllowAllOrigins = true
	}

	router.Use(cors.New(config))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	route.NewCountryRouter(&server.ConcreteServer, router.Group("/countries"))
	route.NewEsimRouter(&server.ConcreteServer, router.Group("/esims"))
	route.NewOrderRouter(&server.ConcreteServer, router.Group("/orders"))

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.Config.ServerPort)
}
