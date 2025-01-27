package main

import (
	config "arno/configs"
	"arno/db"
	handler "arno/internal/handlers"
	"arno/internal/repository"
	"arno/internal/server"
	"log"
)

func main() {
	config.GetDBConfig()

	dbase, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	rep := repository.NewRepository(dbase)
	hnd := handler.NewHandler(rep)
	srv := server.New(hnd)

	err = srv.Run()
	if err != nil {
		log.Println("ошибка при запуске сервера", err)
		return
	}

	defer db.CloseDB(dbase)
}
