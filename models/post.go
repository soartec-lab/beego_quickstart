package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id int64
	Content string `orm: "type(text)"`
	Created time.Time
	Modified time.Time
}

func init() {
	orm.RegisterModel(new(Post))
}

func (this *Post) FindAll() ([]*Post, error) {
	o := orm.NewOrm()
	var posts []*Post
	_, err := o.QueryTable("post").OrderBy("Created").All(&posts)

	return posts, err
}

func (this *Post) FindById(id int64) (*Post, error) {
	o := orm.NewOrm()
	post := Post { Id: id }
	err := o.Read(&post)

	return &post, err
}

func (this *Post) Save(p *Post) error {
	var err error
	o := orm.NewOrm()
	now := time.Now()

	if p.Id != 0 {
		err = o.Read(p)

		if err != nil {
			p.Modified = now
			_, err = o.Update(p)
		}
	} else {
		p.Created = now
		p.Modified = now
		_, err = o.Insert(p)
	}

	return err
}

func (this *Post) Delete(p *Post) error {
    o := orm.NewOrm()
    _, err := o.Delete(p)

    return err
}