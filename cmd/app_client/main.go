package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/lavinas/keel/internal/client/repo"
	"github.com/lavinas/keel/internal/client/core/service"
	"github.com/lavinas/keel/internal/client/hdlr"

)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repo := repo.NewRepoMysql(db)
	create := service.NewCreate(repo)
	list := service.NewList(repo)
	hdlr.NewHandlerChi(create, list)
}