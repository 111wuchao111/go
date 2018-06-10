package models

type Article struct {
	Id            int       `json:"id"`
	User          *User     `json:"user" orm:"rel(fk)"`
	Category      *Category `json:"category" orm:"rel(fk)"`
	Title         string    `json:"tilte"`
	Sub_title     string    `json:"sub_title"`
	Watch_count   int       `json:"watch_count"`
	Comment_conut int       `json:"comment_conut"`
	Good_count    int       `json:"good_count"`
	Is_recommend  int8      `json:"is_recommend"`
	Is_kern       int8      `json:"is_kern"`
	View_img_url  string    `json:"view_img_url"`
	Content       string    `json:"content"`
	Time          string    `json:"time"`
}
