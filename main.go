package main

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/n-jaisabai/q-chang-assignment/infrastructure/database"
	_middleware "github.com/n-jaisabai/q-chang-assignment/middleware"
	"github.com/n-jaisabai/q-chang-assignment/models"
	_problem1Handler "github.com/n-jaisabai/q-chang-assignment/problem_1/delivery/http"
	_problem1Usecase "github.com/n-jaisabai/q-chang-assignment/problem_1/usecase"
	_problem2Handler "github.com/n-jaisabai/q-chang-assignment/problem_2/delivery/http"
	_problem2Repository "github.com/n-jaisabai/q-chang-assignment/problem_2/repository"
	_problem2Usecase "github.com/n-jaisabai/q-chang-assignment/problem_2/usecase"
	"gorm.io/gorm"
)

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error while loading the env:\n %+v", err)
	}

	e := echo.New()

	dbConfig := database.NewConfigPostgres()
	db, err := database.NewSqlHandler(dbConfig)
	if err != nil {
		e.Logger.Fatal(err, "Could not make a connection to the database")
	}
	e.Logger.Infof("Successfully connected to the SQL database")

	if err = db.Conn.AutoMigrate(&models.CashierDesk{}); err == nil && db.Conn.Migrator().HasTable(&models.CashierDesk{}) {
		if err := db.Conn.First(&models.CashierDesk{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			seeds := []*models.CashierDesk{
				{ID: uuid.New(), CoinType: 1000, CoinAmount: 10},
				{ID: uuid.New(), CoinType: 500, CoinAmount: 20},
				{ID: uuid.New(), CoinType: 100, CoinAmount: 15},
				{ID: uuid.New(), CoinType: 50, CoinAmount: 20},
				{ID: uuid.New(), CoinType: 20, CoinAmount: 30},
				{ID: uuid.New(), CoinType: 10, CoinAmount: 20},
				{ID: uuid.New(), CoinType: 5, CoinAmount: 20},
				{ID: uuid.New(), CoinType: 1, CoinAmount: 20},
				{ID: uuid.New(), CoinType: 0.25, CoinAmount: 50},
			}

			for _, seed := range seeds {
				dbTx := db.Conn.Model(&models.CashierDesk{})
				if err := dbTx.Create(seed).Error; err != nil {
					panic("can't seed data")
				}
			}
		}
	}

	// Middleware
	middL := _middleware.InitMiddleware()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middL.CORS)

	timeoutContext := time.Duration(2) * time.Second

	p1u := _problem1Usecase.NewProblem1Usecase()
	_problem1Handler.NewProblem1Handler(e, p1u)

	p2r := _problem2Repository.NewSqlProblem2Repository(db.Conn)
	p2u := _problem2Usecase.NewProblem2Usecase(p2r, timeoutContext)
	_problem2Handler.NewProblem2Handler(e, p2u)

	e.Logger.Fatal(e.Start(":8000"))
}
