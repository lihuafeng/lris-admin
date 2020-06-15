package Controller

import "github.com/kataras/iris"

type IndexController struct {
	base BaseController
}

func (index *IndexController) Index(ctx iris.Context) {
	ctx.View("index.html")
}
