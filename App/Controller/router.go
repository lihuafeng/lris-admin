package Controller

import (
	"github.com/kataras/iris"
)

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("404 ERROR")
}
func RouterHandler(app *iris.Application) {
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)

	app.Get("/", new(IndexController).Index).Name = "home"
}
