package v1

import (
	"github.com/gin-gonic/gin"
	"notification/internal/usecase"
	"notification/pkg/logger"
)

type extensionRoutes struct {
	e usecase.Extension
	l logger.Interface
}

func newExtensionRouter(e *gin.RouterGroup, l logger.Interface, ext usecase.Extension) {
	r := &extensionRoutes{e: ext, l: l}
	g := e.Group("/ext")
	{
		g.POST("/register", r.Register)
	}
}

func (e *extensionRoutes) Register(c *gin.Context) {

}
