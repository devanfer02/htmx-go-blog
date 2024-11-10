package main

import (
	"github.com/devanfer02/go-todo/infra/database"
	"github.com/devanfer02/go-todo/infra/server"
)

func main() {
	pgsqldb := database.NewPgsqlConn()
	httpSrv := server.NewHTTPServer(pgsqldb)

	httpSrv.MountMiddlewares()
	httpSrv.MountControllers()
	httpSrv.Start()
}