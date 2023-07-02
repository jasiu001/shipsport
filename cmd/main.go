package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"shipsport/internal/core/service/shipsport"
	"shipsport/internal/handler"
	"shipsport/internal/repository/shipsportrepo"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	repository := shipsportrepo.NewRepository()
	service := shipsport.NewService(repository)
	hdl := handler.NewFileReader(service)
	if err := hdl.Handle(ctx, os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
