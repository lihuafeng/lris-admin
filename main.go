package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"lris-admin/App/Controller"
	. "lris-admin/App/Middleware"
	"lris-admin/Config"
	"lris-admin/Utils/DB"
	"lris-admin/Utils/Log"
	"lris-admin/Utils/Redis"
	"strconv"
	"time"
)

func initApp() (app *iris.Application) {
	app = iris.New()

	app.Use(recover.New())
	//请求记录
	app.Use(new(AccessMdw).Handler())

	app.RegisterView(iris.Django(Config.VIEW_DIR, Config.VIEW_EXT).Reload(true).Binary(Asset, AssetNames))
	//app.Favicon(Config.ICON_PATH, "/favicon.ico")
	app.StaticEmbedded("/static", "Public", Asset, AssetNames)
	app.StaticWeb("/static", Config.STATIC_DIR)
	//路由实现
	Log.Info("Register Router")
	Controller.RouterHandler(app)

	//连接数据库
	Log.Info("DB CONNECT")
	_ = DB.Connect(Config.DB_DRIVE, Config.DB_DNS)
	defer DB.Close()
	//redis连接池
	Log.Info("DB REDIS")
	_ = Redis.CreatePool(Config.REIDS_ADDR, Config.REDIS_DB, Config.REDIS_PWD)
	defer Redis.ClosePool()
	return
}

func run() {
	app := initApp()
	configuration := iris.WithConfiguration(iris.Configuration{
		TimeFormat: Config.SERVER_TIME_FORMAT,
		Charset:    Config.SERVER_CHARSET,
	})
	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// 关闭所有主机
		_ = app.Shutdown(ctx)
		Log.Info(Config.SERVER_NAME + " Close")
	})
	Log.Info(Config.SERVER_NAME + " Run")
	_ = app.Run(iris.Addr(":"+strconv.Itoa(Config.SERVER_PORT)), configuration)
}

func main() {
	run()
}
