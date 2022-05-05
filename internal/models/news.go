package models

type News struct {
	Id         int
	Header     string
	News       string
	PathToHTML string
}

type PreviewNews struct {
	Id          int
	Header      string
	PathToImage string
}
