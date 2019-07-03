package main

import (
	"com.phh/blog/config"
	"com.phh/blog/web/controller"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

var db *gorm.DB
var cfg *config.Conf

func init() {
	//加载配置
	cfg = config.Cfg()
	//加载数据库配置
	var err error
	db, err = gorm.Open("mysql", "root:root@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
}

func main() {
	//********************创建iris服务**********************
	app := iris.New()
	app.Logger().SetLevel("debug")
	// 加载模版文件
	tmpl := iris.HTML("./web/templates", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)
	//静态资源目录
	app.StaticWeb(cfg.StaticPath, "./web/res")
	//错误处理
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "出错了！"))
		ctx.View("shared/error.html")
	})
	app.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		ctx.ViewData("ctxPath", cfg.ContextPath)
		ctx.View("404.html")
	})
	app.Use(recover.New())
	app.Use(func(ctx context.Context) {
		fmt.Println(">>>>>>:" + cfg.ContextPath + " ,url:" + ctx.RequestPath(false))
		ctx.ViewData("ctxPath", cfg.ContextPath)
		ctx.Next()
	})
	//app.UseGlobal() //此方法会拦截所有请求包括静态资源，而app.Use静态资源不会
	//会话管理
	sess := sessions.New(sessions.Config{Cookie: "GOSESSIONID"})
	//root context-path
	root := mvc.New(app.Party(cfg.ContextPath))
	{
		root.Register(sess.Start)
		//register index
		root.Handle(new(controller.IndexController))
		//register users
	}
	//启动应用
	app.Run(
		// 启动端口
		iris.Addr(cfg.ServerAddress),
		//iris.WithoutVersionChecker,
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// 优化
		iris.WithOptimizations,
	)
}
