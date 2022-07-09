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
		g.POST("/register", r.register)
		//g.Any("/connect", r.Connect)
		g.GET("/:uid/:group", r.getGroup)
		g.POST("/:uid", r.addGroup)
		g.DELETE("/:uid/:group", r.removeGroup)
	}
}

// register godoc
// @Summery      register extension client
// @Description  get uid
// @Tags         extension
// @Param        name       query  string  true  "client name"   minlength(1)  maxlength(64)
// @Param        extension  query  string  true  "extension id"  minlength(1)  maxlength(256)
// @Produce      json
// @Success      200  {object}  ExtensionResp{data=object{uid=string}}
// @Router       /v1/ext/register [post]
func (e *extensionRoutes) register(c *gin.Context) {
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

func (e *extensionRoutes) addGroup(c *gin.Context) {

}

func (e *extensionRoutes) getGroup(c *gin.Context) {

}

func (e *extensionRoutes) removeGroup(c *gin.Context) {

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
