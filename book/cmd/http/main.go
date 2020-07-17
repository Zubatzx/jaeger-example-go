package main

import (
	"book/internal/boot"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := boot.HTTP(); err != nil {
		log.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
}
