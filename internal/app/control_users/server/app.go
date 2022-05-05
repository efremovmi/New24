package server

import (
	ctrlUsers "News24/internal/app/control_users"
	endPoints "News24/internal/app/control_users/delivery/http"
	"News24/internal/app/control_users/repository/postgres"
	"News24/internal/app/control_users/usecase"
	"News24/internal/models"

	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	httpServer       *http.Server
	config           *models.Config
	ctrlUsersUseCase ctrlUsers.UseCase
}

func NewApp(config *models.Config) (app *App) {
	db, err := initDB(config)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ctrlUsersService := usecase.NewContolUserUseCase(db, fmt.Sprintf("%v", config.HASH_SALT))

	app = &App{
		ctrlUsersUseCase: ctrlUsersService,
		config:           config}

	return app
}

func (a *App) Run() {
	router := gin.Default()

	endPoints.RegisterHTTPEndpoints(router, a.ctrlUsersUseCase)

	router.Run(fmt.Sprintf("%v", a.config.ADR_CONTROL_USERS))
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
