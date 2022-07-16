package v1

import (
	"aio/internal/entity"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type barkRoutes struct {
	b usecase.Bark
	l logger.Interface
}

func newBarkRouter(g *gin.RouterGroup, l logger.Interface, b usecase.Bark) {
	e := g.Group("/bark")
	r := &barkRoutes{b, l}
	e.GET("/ping", r.ping)
	e.GET("/health", r.health)
	e.GET("/info", r.info)
	e.GET("/register", r.register)
	e.GET("/:key/:content", r.push)
	e.GET("/:key", r.pull)
}

func (r *barkRoutes) ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":      200,
		"message":   "pong",
		"timestamp": time.Now().Unix(),
	})
}

func (r *barkRoutes) health(c *gin.Context) {
	c.String(200, "ok")
}

func (r *barkRoutes) info(c *gin.Context) {
	c.JSON(200, gin.H{
		// TODO: add more info
		"version": 0,
		"build":   "",
		"arch":    "",
		"commit":  "",
		"devices": []string{},
	})
}

func (r *barkRoutes) register(c *gin.Context) {
	token := c.Query("devicetoken")
	key := c.Query("key")
	name := c.DefaultQuery("name", "none")

	ent := &entity.BarkDevice{
		DeviceToken: token,
		DeviceKey:   key,
		Name:        name,
	}
	err := r.b.Register(c.Request.Context(), ent)
	if err != nil {
		r.l.Errorln("register bark device", err)
		barkResp(c, 400, "internal error", nil)
		return
	}
	barkResp(c, 200, "success", gin.H{
		"key":          ent.DeviceKey,
		"device_key":   ent.DeviceKey,
		"device_token": ent.DeviceToken,
	})
}

func (r *barkRoutes) push(c *gin.Context) {
	key := c.Param("key")

	m := make(map[string]string)
	for k, v := range c.Request.URL.Query() {
		if len(v) != 0 {
			m[k] = v[0]
		}
	}

	msg := &entity.APNsMessage{
		Title:    c.Query("title"),
		Category: c.Query("category"),
		Body:     c.Param("content"),
		Sound:    c.DefaultQuery("sound", "1107"),
		Data:     m,
	}

	err := r.b.Push(c.Request.Context(), key, msg)
	if err != nil {
		r.l.Errorln("push bark message", err)
		barkResp(c, 400, "internal error", nil)
		return
	}
	barkResp(c, 200, "success", nil)
}

func (r *barkRoutes) pull(c *gin.Context) {
	key := c.Param("key")
	list, err := r.b.Pull(nil, key, 0, 100)
	if err != nil {
		r.l.Errorln("fetch bark history", err)
		barkResp(c, 400, "internal error", nil)
		return
	}
	barkResp(c, 200, "success", list)
}
