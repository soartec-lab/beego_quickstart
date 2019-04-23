package controllers

import (
	"github.com/astaxie/beego"
	"github.com/soartec-lab/beego_quickstart/models"
)

type PostController struct {
	beego.Controller
	post models.Post
}

func (this *PostController) Index() {
	posts, _ := this.post.FindAll()
	this.Data["posts"] = posts
	this.Layout = "layouts/application.html"
	this.TplName = "post/index.html"
}