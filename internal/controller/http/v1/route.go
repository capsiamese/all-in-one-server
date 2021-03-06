package v1

import (
	"aio/docs"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type UseCase struct {
	B usecase.Bark
	P usecase.PushDeer
	E *usecase.ExtensionUseCase
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

func NewRouter(e *gin.Engine, l logger.Interface, u *UseCase) {
	docs.SwaggerInfo.Title = "Hello"

	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	e.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	e.GET(StatsVis())

	h := e.Group("/v1")
	{
		newBarkRouter(h, l, u.B)
		newPushDeerRouter(h, l, u.P)
		newExtensionRouter(h, l, u.E)
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func StatsVis() (string, gin.HandlerFunc) {
	const statsVisPath = "/debug/statsviz/*filepath"
	return statsVisPath, func(c *gin.Context) {
		if c.Param("filepath") == "/ws" {
			statsviz.Ws(c.Writer, c.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(c.Writer, c.Request)
	}
}
