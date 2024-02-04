package api

import (
	"esim-service/api/route"
	"esim-service/config"
	db "esim-service/db/sqlc"
	"esim-service/domain"
	"esim-service/domain/order"
	repository "esim-service/repository/http"
	"esim-service/service"
	"esim-service/usecase"
	"net/http"
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

	server := &Server{
		ConcreteServer: domain.ConcreteServer{Store: store, Config: env, Timeout: timeout},
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

	xhr := repository.NewXenditHttpRepository(server.Config)
	uhr := repository.NewUsimsaHttpRepository(server.Config)
	ss := service.NewSmtpService(server.Config)
	ou := usecase.NewOrderUsecase(&server.ConcreteServer, ss, xhr, uhr)
	router.POST("/webhook", webhookHandler(ou))

	route.NewCountryRouter(&server.ConcreteServer, router.Group("/countries"))
	route.NewEsimRouter(&server.ConcreteServer, router.Group("/esims"))
	route.NewOrderRouter(&server.ConcreteServer, router.Group("/orders"))

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":" + server.Config.ServerPort)
}

func webhookHandler(ou order.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			var req order.UsimsaSubscribedOrderReq

			err := c.ShouldBindJSON(&req)
			if err != nil {
				c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
				return
			}

			order, err := ou.SendPurchasedEsimEmail(c, req.TopupId, req.Iccid, req.Smdp, req.ActivateCode, req.QrcodeImgUrl)
			if err != nil {
				c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
				return
			}

			c.JSON(http.StatusOK, order)

		} else {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"message": "Invalid request method",
			})
		}
	}
}
