package main

import (
	"TodoList/conf"
	"TodoList/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conf.Init()
	routers.NewRouters()
}
