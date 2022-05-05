package news

import (
	"News24/internal/models"
)

type NewsRepository interface {
	SaveNews(news *models.News) (err error)
	DeleteNewsForHeader(header string) (err error)
	UpdateNewsForId(newNews *models.News, id int) (err error)
	GetNewsForHeader(header string) (news *models.News, err error)
	GetListPreviewNews(lastId int) (previewNewsList []*models.PreviewNews, err error)
}
