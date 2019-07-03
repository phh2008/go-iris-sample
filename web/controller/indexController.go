package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

type IndexController struct {
	Session *sessions.Session
}

//localhost:8080/
func (c *IndexController) Get() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":  "首页",
			"navIndex": "layui-this",
		},
	}
}

//localhost:8080/about
func (c *IndexController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "about.html",
		Data: iris.Map{
			"Title":  "关于",
			"navAbout": "layui-this",
		},
	}
}
