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

func (this *PostController) Create() {
	post := models.Post {
		Content: this.GetString("content"),
	}
	err := this.post.Save(&post)
	flash := beego.NewFlash()

	if err != nil {
		flash.Error("The post could not be saved. Please, try again.")
	} else {
		flash.Notice("The post has been saved.")
	}

	flash.Store(&this.Controller)

	this.Redirect("/", 302)
}