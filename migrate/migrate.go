package main

import (
	"github.com/pulse227/server-recruit-challenge-sample/db"
	"github.com/pulse227/server-recruit-challenge-sample/model"
)

func main()  {
	dbConn := db.NewDB()
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Singer{}, &model.Album{})
}