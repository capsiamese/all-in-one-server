package v1

import (
	"github.com/gin-gonic/gin"
	"notification/pkg/logger"
	"runtime"
	"strings"
	"time"
)

func barkResp(c *gin.Context, code int, msg string, data interface{}) {
	h := gin.H{
		"code":      code,
		"message":   msg,
		"timestamp": time.Now().Unix(),
	}
	if data != nil {
		h["data"] = data
	}
	c.JSON(code, h)
}

func pdResp(c *gin.Context, state, code int, err error, content interface{}) {
	m := gin.H{
		"code": code,
	}
	if err != nil {
		m["error"] = err
	} else if content != nil {
		m["content"] = content
	}
	c.JSON(state, m)
}

type ExtensionResp struct {
	Code int   `json:"code"`
	Err  error `json:"error,omitempty" swaggertype:"string"`
	Data any   `json:"data,omitempty"`
} //@name ExtensionResponse

func ExtResp(c *gin.Context, state, code int, err error, data any) {
	if err != nil {
		ctx := c.Request.Context()
		if l := ctx.Value("logger"); l != nil {
			if log, ok := l.(logger.Interface); ok {
				var rootDir string
				if dir := ctx.Value("rootDir"); dir != nil {
					if d, ok := dir.(string); ok {
						rootDir = d
					}
				}
				_, fn, ln, _ := runtime.Caller(1)
				fn = strings.TrimPrefix(fn, rootDir)
				log.Errorf("file:%s line:%d %v", fn, ln, err)
			}
		}
	}
	c.JSON(state, &ExtensionResp{
		Code: code,
		Err:  err,
		Data: data,
	})
}
