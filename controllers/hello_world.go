package controllers

import (
	"fmt"
	"net/http"

	"github.com/srfbogomolov/repository-db/sqldb"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if err := sqldb.DB.Ping(); err != nil {
		fmt.Println("DB Error")
	}

	w.Write([]byte("Hello, world"))
}
