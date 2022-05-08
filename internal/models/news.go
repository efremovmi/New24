package models

type News struct {
	Id         int    `json:"id"`
	Header     string `json:"header"`
	News       string `json:"news"`
	PathToHTML string `json:"-"`
}

type PreviewNews struct {
	Id          int
	Header      string
	PathToImage string
}
