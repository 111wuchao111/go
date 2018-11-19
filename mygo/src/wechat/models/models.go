package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ArticleClass struct {
	Id        int    `json:"id"`
	ClassName string `json:"class_name"`
}

func init() {
	orm.RegisterModel(new(ArticleClass))
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/wechatapp?charset=utf8")
}
