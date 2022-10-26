package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sleekflow/app/api/middleware"

	"time"

	"gorm.io/driver/postgres"

	"sleekflow/app/api/server"
	"sleekflow/lib/database_transaction"
	"sleekflow/lib/validators"

	toDoHTTP "sleekflow/service/to_do/delivery"
	toDoModule "sleekflow/service/to_do/module"
	usersHTTP "sleekflow/service/users/delivery"
	usersModule "sleekflow/service/users/module"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/subosito/gotenv"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type libs struct {
	fx.Out

	TransactionManager database_transaction.Client
}

type handlers struct {
	fx.In

	ToDoHandler  *toDoHTTP.Handler
	UsersHandler *usersHTTP.Handler
}

func main() {
	log.Println("server is starting")

	loadEnv()

	// set log to file
	if os.Getenv("APP_ENV") != "development" {
		log.Println("running in ", os.Getenv("APP_ENV"), " environment")
		f, err := os.OpenFile("error-log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//defer to close when you're done with it, not because you think it's idiomatic!
		defer f.Close()

		//set output of logs to f
		log.SetOutput(f)
	}

	app := fx.New(
		fx.Provide(
			setupDatabase,
			initLibs,
		),
		toDoModule.Module,
		usersModule.Module,
		fx.Invoke(
			validators.NewValidator,
			startServer,
		),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, handlers handlers) {
	m := middleware.New(
		middleware.Config{},
	)

	h := server.BuildHandler(m,
		handlers.ToDoHandler,
		handlers.UsersHandler,
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      h,
		ReadTimeout:  300 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func(s *http.Server) {
				log.Printf("api is available at %s\n", s.Addr)
				if err := s.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}(s)
			return nil
		},
		OnStop: func(c context.Context) error {
			_ = s.Shutdown(c)
			log.Println("api gracefully stopped")
			return nil
		},
	})
}

func loadEnv() {
	err := gotenv.Load()

	if err != nil {
		log.Println("failed to load from .env")
	}
}

func setupDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return db
}

func initLibs(lc fx.Lifecycle, db *gorm.DB) libs {
	l := libs{
		TransactionManager: database_transaction.New(db),
	}

	return l
}
