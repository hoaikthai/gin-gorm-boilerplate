package main

import (
	"flag"
	"fmt"
	"gin-gorm-boilerplate/config"
	"gin-gorm-boilerplate/db"
	"gin-gorm-boilerplate/server"
	"os"
)

func main() {
	env := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	config.Init(*env)
	db.Init()
	server.Init()
}
