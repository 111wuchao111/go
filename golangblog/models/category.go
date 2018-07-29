package models

type Category struct {
	Id      int        `json:"id"`
	Name    string     `json:"name"`
	Article []*Article `orm:"reverse(many)"`
}

// 自定义表名（系统自动调用）
func (c *Category) TableName() string {
	return "article_category"
}
