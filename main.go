package main

import (
	"github.com/guilherme-brandao/to-go-list/database"
	"github.com/guilherme-brandao/to-go-list/server"
)

func main() {
	database.Init()
	server.Init()
}