package server

import (
	"News24/internal/app/news"
	endPoints "News24/internal/app/news/delivery/http"
	"News24/internal/app/news/repository/postgres"
	"News24/internal/app/news/usecase"
	"News24/internal/models"
	"fmt"
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type App struct {
	httpServer  *http.Server
	config      *models.Config
	newsUseCase news.UseCase
}

func NewApp(config *models.Config) (app *App) {
	db, err := initDB(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = os.Mkdir("storage", 0777)
	if err != nil {
		if err.Error() != "mkdir storage: file exists" {
			log.Fatalf("Error: %v", err)
		}
	}

	newsService := usecase.NewNewsUseCase(db, fmt.Sprintf("%v", "views"))

	app = &App{
		config:      config,
		newsUseCase: newsService}

	return app
}

func (a *App) Run() {
	router := gin.Default()
	router.HTMLRender = gintemplate.Default()

	endPoints.RegisterHTTPEndpoints(router, a.newsUseCase)

	router.Run(fmt.Sprintf("%v", a.config.ADR_NEWS))
}

func initDB(config *models.Config) (db *postgres.NewsRepository, err error) {
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_BD_NAME)

	db = &postgres.NewsRepository{}
	err = db.NewUserRepository(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_NEWS))

	return db, err
}
