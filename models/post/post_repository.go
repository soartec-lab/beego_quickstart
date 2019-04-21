package post

import (
	"github.com/astaxie/beego/orm"
)

type PostRepository struct {}

func init() {
	orm.RegisterModel(new(Post))
}