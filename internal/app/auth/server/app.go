package server

import (
	"News24/internal/app/auth"
	endPoints "News24/internal/app/auth/delivery/http"
	"News24/internal/app/auth/repository/postgres"
	"News24/internal/app/auth/usecase"
	"News24/internal/models"
	gintemplate "github.com/foolin/gin-template"

	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type App struct {
	httpServer  *http.Server
	config      *models.Config
	authUseCase auth.UseCase
}

func NewApp(config *models.Config) (app *App) {
	db, err := initDB(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	authService := usecase.NewAuthUseCase(db, fmt.Sprintf("%v", config.HASH_SALT), time.Duration(5))

	app = &App{
		authUseCase: authService,
		config:      config}

	return app
}

func (a *App) Run() {
	router := gin.Default()
	router.HTMLRender = gintemplate.Default()
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("/home/umd/gop/testcss/static"))))

	endPoints.RegisterHTTPEndpoints(router, a.authUseCase, fmt.Sprintf("%v", a.config.PATH_TO_VIEWS))

	router.Run(fmt.Sprintf("%v", a.config.ADR_AUTH))
}

func initDB(config *models.Config) (db *postgres.UserRepository, err error) {
	psqlconn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_BD_NAME)

	db = &postgres.UserRepository{}
	err = db.NewUserRepository(psqlconn, fmt.Sprintf("%v", config.POSTGRES_TABLE_USERS))

	return db, err
}
