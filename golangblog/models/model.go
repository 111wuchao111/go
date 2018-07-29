package models

import (
	"crypto/md5"
	"encoding/hex"

	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterModel(new(Brand), new(Article), new(User), new(Category))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/goblog?charset=utf8")
	orm.Debug = true
}

//参数int/string
func GetMd5(v interface{}) string {
	var s string
	if v2, ok := v.(string); ok {
		s = v2
	} else {
		s = strconv.Itoa(v.(int))
	}
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
