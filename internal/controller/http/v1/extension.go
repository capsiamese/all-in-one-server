package v1

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
		g.Any("/connect", r.Connect)
	}
}

func (e *extensionRoutes) Register(c *gin.Context) {
	name := c.Query("name")
	extension := c.Query("extension")
	obj, err := e.e.Register(c.Request.Context(), name, extension)
	if err != nil {
		ExtResp(c, 200, 1, err, nil)
		return
	}
	ExtResp(c, 200, 0, nil, gin.H{
		"uid": obj.ClientUID,
	})
}

func (e *extensionRoutes) Connect(c *gin.Context) {
	uid, err := uuid.FromString(c.Query("uid"))
	if err != nil {
		ExtResp(c, 200, 0, err, nil)
		return
	}
	err = e.e.Connect(c.Request.Context(), uid, nil)
	if err != nil {
		ExtResp(c, 200, 0, err, nil)
		return
	}
}
