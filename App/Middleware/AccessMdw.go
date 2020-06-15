package Middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"lris-admin/Utils/Log"
	"strings"
	"time"
)

var skipExt = []string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

type AccessMdw struct {
}

func (mdw AccessMdw) Handler() context.Handler {
	conf := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}

	conf.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		Log.Access(output)
	}
	conf.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range skipExt {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	handler := logger.New(conf)
	return handler
}
