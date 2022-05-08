package news

import (
	"News24/internal/models"
	"mime/multipart"
)

type UseCase interface {
	SaveNews(image multipart.File, header, news, imageFileName string) (statusCode int, err error)
	DeleteNewsForHeader(header string) (statusCode int, err error)
	UpdateNewsForId(image multipart.File, header, news, imageFileName string, id int) (statusCode int, err error)
	GetNewsHTMLForHeader(header string) (news *models.News, statusCode int, err error)
	GetListPreviewNews(lastId int) (previewNewsList []*models.PreviewNews, statusCode int, err error)
	GetAllNews() (allNews []*models.News, pathToImage string, statusCode int, err error)
}
