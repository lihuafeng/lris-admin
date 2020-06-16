package Admin

import (
	"github.com/kataras/iris"
)

type IndexController struct{}

func (index *IndexController) Index(ctx iris.Context) {
	ctx.ViewData("Abc", "abc")
	ctx.View("admin/index.html")
}

func (index *IndexController) Desktop(ctx iris.Context) {
	ctx.View("admin/desktop.html")
}
