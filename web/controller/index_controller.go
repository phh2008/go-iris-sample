package controller

import (
	"com.phh/blog/models"
	"com.phh/blog/services"
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/rs/zerolog/log"
)

type IndexController struct {
	Session        *sessions.Session
	ArticleService services.ArticleService
}

//localhost:8080/
func (c *IndexController) Get() mvc.Result {
	articles := c.ArticleService.Query(models.Article{})
	json, _ := json.Marshal(articles)
	log.Info().Msg(string(json))
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "首页",
			"navIndex": "layui-this",
			"articles": articles,
		},
	}
}

//localhost:8080/about
func (c *IndexController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "about.html",
		Data: iris.Map{
			"Title":    "关于",
			"navAbout": "layui-this",
		},
	}
}
