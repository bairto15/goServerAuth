package main

import (
	"goServerAuth/package/handler"
	"goServerAuth/package/repository"
	"goServerAuth/package/service"
	"goServerAuth/server"

	"github.com/gonutz/w32"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	//Логи в режиме JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//при запуске запускается в фоном режиме
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}

	db, err := repository.NewPostgersDB(repository.Store{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "123789",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("Error running PostgerSQL: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("14625", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error running http Server: %s", err.Error())
	}
}
