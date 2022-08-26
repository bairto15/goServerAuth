package main

import (
	"github.com/gonutz/w32"
	"github.com/sirupsen/logrus"
	"goServerAuth/package/handler"
	"goServerAuth/package/repository"
	"goServerAuth/package/service"
	"goServerAuth/server"
)

func main() {
	//Логи в режиме JSON
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//при запуске build запускается в фоном режиме
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, w32.SW_HIDE)
		}
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("14625", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error running http Server: %s", err.Error())
	}
}
