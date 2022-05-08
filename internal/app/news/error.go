package news

import "errors"

var (
	NewsNotFound             = errors.New("Новость не найдена")
	BDNotWorking             = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD = errors.New("Error in database connection parameters")
	BadInsertNews            = errors.New("Failed to insert news")
	BadUpdateNews            = errors.New("Failed update news")
	BadDeleteNews            = errors.New("Failed delete news")
	BadGetPreviewNewsList    = errors.New("Ошибка получения списка превью новостей")
	BadGetNewsList           = errors.New("Ошибка получения списка всех новостей")
	BadImage                 = errors.New("Ошибка при сохранении файла, перепроверьте изображение")
	NewsFoundDuplicate       = errors.New("Пост с такой шапкой уже есть, придумайте другой")
	ServerBDError            = errors.New("Внутренняя ошибка при работе с базой данных")
	BadInputFields           = errors.New("Не указана шапка или тело новости")
	ServerError              = errors.New("Ошибка на стороне сервера")
	BadRequest               = errors.New("Некорректное тело запроса")
	ElementsNotFound         = errors.New("Элементы не найдены")
	BadId                    = errors.New("Некорректный id")
)
