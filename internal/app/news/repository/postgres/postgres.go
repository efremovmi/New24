package postgres

import (
	errorsCustom "News24/internal/app/news"
	"News24/internal/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type News struct {
	ID          int    `json:"id,omitempty"`
	Header      string `json:"username"`
	NewsText    string `json:"password"`
	PathToImage int    `json:"role"`
}

type NewsRepository struct {
	psqlconn  string
	tableName string
}

func (r *NewsRepository) NewUserRepository(psqlconn, tableName string) (err error) {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return errorsCustom.BDNotWorking
	}
	r.tableName = tableName
	r.psqlconn = psqlconn
	return nil
}

func (r *NewsRepository) SaveNews(news *models.News) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO %s (header, news_text, path_to_html) VALUES ('%s', '%s', '%s') RETURNING id;",
		r.tableName,
		news.Header,
		news.News,
		news.PathToHTML)
	var id int
	if err = db.QueryRow(query).Scan(&id); err != nil {
		return errorsCustom.BadInsertNews
	}
	return nil
}

func (r *NewsRepository) DeleteNewsForHeader(header string) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("delete from %s where header = '%s' returning id;", r.tableName, header)

	var id int
	if err = db.QueryRow(query).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return errorsCustom.NewsNotFound
		}
		return errorsCustom.BadDeleteNews
	}
	return nil
}

func (r *NewsRepository) UpdateNewsForId(newNews *models.News, id int) (err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("UPDATE %s set news_text = '%s', header = '%s', path_to_html = '%s' where id = %d returning id;",
		r.tableName,
		newNews.News,
		newNews.Header,
		newNews.PathToHTML,
		id)

	if err = db.QueryRow(query).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return errorsCustom.NewsNotFound
		}
		return errorsCustom.BadUpdateNews
	}
	return nil
}

func (r *NewsRepository) GetNewsForHeader(header string) (news *models.News, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return nil, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, header, news_text, path_to_html FROM %s WHERE header = '%s';",
		r.tableName,
		header)

	news = &models.News{}

	if err = db.QueryRow(query).Scan(&news.Id, &news.Header, &news.News, &news.PathToHTML); err != nil {
		return nil, errorsCustom.NewsNotFound
	}

	return news, nil
}

func (r *NewsRepository) GetNewsForId(id int) (news *models.News, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	if err != nil {
		return nil, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id, header, news_text, path_to_html FROM %s WHERE id = %d;",
		r.tableName,
		id)

	news = &models.News{}

	if err = db.QueryRow(query).Scan(&news.Id, &news.Header, &news.News, &news.PathToHTML); err != nil {
		return nil, errorsCustom.NewsNotFound
	}

	return news, nil
}

func (r *NewsRepository) GetListPreviewNews(lastId int) (previewNewsList []*models.PreviewNews, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	previewNewsList = make([]*models.PreviewNews, 0)

	if err != nil {
		return previewNewsList, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("select id, header, path_to_html from %s where id > %d limit 5;", r.tableName, lastId)

	rows, err := db.Query(query)
	if err != nil {
		return make([]*models.PreviewNews, 0), errorsCustom.BadGetPreviewNewsList
	}

	for rows.Next() {
		previewNews := models.PreviewNews{}
		err = rows.Scan(&previewNews.Id, &previewNews.Header, &previewNews.PathToImage)
		if err != nil {
			return make([]*models.PreviewNews, 0), errorsCustom.BadGetPreviewNewsList
		}

		previewNewsList = append(previewNewsList, &previewNews)
	}

	return previewNewsList, nil
}

func (r *NewsRepository) GetAllNews() (newsList []*models.News, err error) {
	db, err := sql.Open("postgres", r.psqlconn)
	newsList = make([]*models.News, 0)

	if err != nil {
		return newsList, errorsCustom.IncorrectParamsConnectBD
	}
	defer db.Close()

	query := fmt.Sprintf("select id, header, news_text from %s;", r.tableName)

	rows, err := db.Query(query)
	if err != nil {
		return make([]*models.News, 0), errorsCustom.BadGetNewsList
	}

	for rows.Next() {
		news := models.News{}
		err = rows.Scan(&news.Id, &news.Header, &news.News)
		if err != nil {
			return make([]*models.News, 0), errorsCustom.BadGetNewsList
		}

		newsList = append(newsList, &news)
	}

	return newsList, nil
}
