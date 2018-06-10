package models

type Brand struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Sub_title string `json:"sub_title"`
	Img_url   string `json:"img_url"`
	Href      string `json:"href"`
}
