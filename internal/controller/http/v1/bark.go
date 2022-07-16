package v1

import (
	"aio/internal/entity"
	"aio/internal/usecase"
	"aio/pkg/logger"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"strings"
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
	e.POST("/:key/:content", r.push)
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
	name := c.DefaultQuery("name", gofakeit.PetName())

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
	var msg = &entity.APNsMessage{}
	if strings.HasPrefix(c.ContentType(), gin.MIMEJSON) {
		err := c.ShouldBindJSON(msg)
		if err != nil {
			r.l.Errorln("push bark message", err)
			barkResp(c, 400, "internal error", nil)
			return
		}
	} else {
		m := make(map[string]string)
		for k, v := range c.Request.URL.Query() {
			if len(v) != 0 {
				m[k] = v[0]
			}
		}
		msg.Title = c.Query("title")
		msg.Category = c.Query("category")
		msg.Content = c.Param("content")
		msg.Sound = c.DefaultQuery("sound", "1107")
		msg.Params = m
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
	params := new(entity.HistoryParam)
	if err := c.ShouldBindQuery(params); err != nil {
		r.l.Errorln("fetch bark history", err)
		barkResp(c, 400, "internal error", nil)
		return
	}

	list, err := r.b.Pull(c.Request.Context(), key, params.Offset, params.Limit)
	if err != nil {
		r.l.Errorln("fetch bark history", err)
		barkResp(c, 400, "internal error", nil)
		return
	}
	r.l.Debugln(list)
	barkResp(c, 200, "success", list)
}
