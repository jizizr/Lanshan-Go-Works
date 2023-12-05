package main

import (
	"ezgin/boot"
	"ezgin/dao/database"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		database.UserDb.SaveToFile("USER")
		database.CommentDb.SaveToFile("COMMENT")
		os.Exit(0)
	}()
	boot.InitRouters()
}
