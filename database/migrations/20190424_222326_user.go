package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20190424_222326 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20190424_222326{}
	m.Created = "20190424_222326"

	migration.Register("User_20190424_222326", m)
}

// Run the migrations
func (m *User_20190424_222326) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20190424_222326) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
