package usecase

import (
	"News24/internal/app/news"
	errorsCustom "News24/internal/app/news"
	"News24/internal/models"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type NewsUseCase struct {
	newsRepo      news.NewsRepository
	pathToStorage string
}

func NewNewsUseCase(newsRepo news.NewsRepository, pathToStorage string) *NewsUseCase {
	return &NewsUseCase{
		newsRepo:      newsRepo,
		pathToStorage: pathToStorage,
	}
}

func (n *NewsUseCase) SaveNews(image multipart.File, header, newsText, imageFileName string) (statusCode int, err error) {
	if len(header) == 0 || len(newsText) == 0 {
		return http.StatusBadRequest, errorsCustom.BadInputFields
	}
	if len(strings.Split(imageFileName, ".jpeg")) != 2 && len(strings.Split(imageFileName, ".png")) != 2 {
		return http.StatusBadRequest, errorsCustom.BadImage
	}

	newsModels, err := n.newsRepo.GetNewsForHeader(header)
	if err == errorsCustom.IncorrectParamsConnectBD {
		return http.StatusInternalServerError, errorsCustom.ServerBDError
	}
	if newsModels != nil && err == nil {
		return http.StatusBadRequest, errorsCustom.NewsFoundDuplicate
	}

	pathToNewNewsFolder := n.pathToStorage + "/" + header
	err = os.Mkdir(pathToNewNewsFolder, 0777)
	if err != nil {
		a := err.Error()
		_ = a
		if err.Error() == "mkdir "+n.pathToStorage+"/"+header+": file exists" {
			return http.StatusBadRequest, errorsCustom.NewsFoundDuplicate
		}
		return http.StatusInternalServerError, errorsCustom.ServerError
	}

	relPathToImage := pathToNewNewsFolder + "/" + header + ".jpeg"
	imageOnServer, err := os.Create(relPathToImage)

	if err != nil {
		return http.StatusInternalServerError, errorsCustom.ServerError
	}
	defer imageOnServer.Close()
	_, err = io.Copy(imageOnServer, image)
	if err != nil {
		return http.StatusBadRequest, errorsCustom.BadImage
	}

	templateHTML := fmt.Sprintf("<!DOCTYPE html>\n"+
		"<html lang=\"en\">\n"+
		"<head>\n"+
		"  <title>News</title>\n"+
		"  <meta charset=\"UTF-8\">\n"+
		"  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">"+
		"\n"+
		"\n"+
		"\n"+
		"\n"+
		"\n"+
		"  <link rel=\"stylesheet\" type=\"text/css\" href=./views/static/css/style.css>\n"+
		"\n"+
		"\n"+
		"</head>\n"+
		"<body>\n"+
		" <form action=\"\" method=\"post\" onsubmit=\"\" id=\"form\">\n"+
		"    <div class=figure>\n"+
		"      <p class=\"head\"> %s"+
		"      <p><img class=scaled src=\"%s\"\n"+
		"        alt=\"Picture\">\n"+
		"      <div class=\"main_text\">%s</div>\n"+
		"     \n"+
		"    </div>\n"+
		"	 <div class=\"menu\">\n"+
		"     <div class=\"menu_text\">#MY #YOUR #NEWS</div>\n"+
		"      <button type=\"submit\">\n"+
		"        Выйти\n"+
		"      </button>\n"+
		"    </div>"+
		"  </form>\n"+
		"<script src=\"./views/static/js/exit.js\"></script>\n"+
		"</body>\n"+
		"</html>\n", header, "./views/"+header+"/"+header+".jpeg", newsText)

	relPathToHTML := pathToNewNewsFolder + "/" + header + ".html"
	fileHTMl, err := os.Create(relPathToHTML)

	if err != nil {
		return http.StatusInternalServerError, errorsCustom.ServerError
	}

	absPathToHTML, err := filepath.Abs(relPathToHTML)
	if err != nil {
		return http.StatusInternalServerError, errorsCustom.ServerError
	}
	defer fileHTMl.Close()
	fileHTMl.WriteString(templateHTML)

	err = n.newsRepo.SaveNews(&models.News{
		Header:     header,
		News:       newsText,
		PathToHTML: absPathToHTML,
	})

	if err != nil {
		os.RemoveAll(pathToNewNewsFolder)
		return http.StatusInternalServerError, errorsCustom.ServerBDError
	}
	return http.StatusOK, nil
}

func (n *NewsUseCase) DeleteNewsForHeader(header string) (statusCode int, err error) {
	if len(header) == 0 {
		return http.StatusBadRequest, errorsCustom.BadInputFields
	}

	newsModel, err := n.newsRepo.GetNewsForHeader(header)
	if err != nil {
		if err == errorsCustom.NewsNotFound {
			return http.StatusBadRequest, errorsCustom.NewsNotFound
		}
		return http.StatusInternalServerError, errorsCustom.ServerBDError
	}

	pathToFolderWithHTML := strings.Split(newsModel.PathToHTML, header)[0] + header

	err = os.RemoveAll(pathToFolderWithHTML)
	if err != nil {
		return http.StatusInternalServerError, errorsCustom.ServerError
	}

	err = n.newsRepo.DeleteNewsForHeader(header)
	if err != nil {
		if err == errorsCustom.NewsNotFound {
			return http.StatusBadRequest, errorsCustom.NewsNotFound
		}
		return http.StatusInternalServerError, errorsCustom.ServerBDError
	}

	return http.StatusOK, nil
}

func (n *NewsUseCase) UpdateNewsForId(newNews models.News, id int) (err error) {
	return nil
}

func (n *NewsUseCase) GetNewsHTMLForHeader(header string) (news *models.News, statusCode int, err error) {
	if len(header) == 0 {
		return nil, http.StatusBadRequest, errorsCustom.BadInputFields
	}
	newsModel, err := n.newsRepo.GetNewsForHeader(header)
	if err != nil {
		if err == errorsCustom.NewsNotFound {
			return nil, http.StatusBadRequest, errorsCustom.NewsNotFound
		}
		return nil, http.StatusInternalServerError, errorsCustom.ServerBDError
	}
	return newsModel, http.StatusOK, nil
}

func (n *NewsUseCase) GetListPreviewNews(lastId int) (previewNewsList []*models.PreviewNews, statusCode int, err error) {
	previewNewsList, err = n.newsRepo.GetListPreviewNews(lastId)
	if err != nil {
		if err == errorsCustom.BadGetPreviewNewsList {
			return previewNewsList, http.StatusOK, nil
		}
		return previewNewsList, http.StatusInternalServerError, errorsCustom.ServerBDError
	}
	return previewNewsList, http.StatusOK, nil
}
