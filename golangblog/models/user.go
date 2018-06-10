package models

type User struct {
	Id               int        `json:"id"`
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	Avatar           string     `json:"avatar"`
	Title            string     `json:"title"`
	Motto            string     `json:"motto"`
	Intro            string     `json:"intro"`
	Desc             string     `json:"desc"`
	Phone            string     `json:"phone"`
	Qq               string     `json:"qq"`
	Email            string     `json:"email"`
	Wx_number        string     `json:"wx_number"`
	Wx_url           string     `json:"wx_url"`
	Wx_money_url     string     `json:"wx_money_url"`
	Alipay_money_url string     `json:"alipay_money_url"`
	Time             string     `json:"time"`
	Artice           []*Article `orm:"reverse(many)"`
}
