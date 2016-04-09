package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"gitlab.fg/go/logger"
)

func getLogger(filepath string) *logger.ServiceLogger {
	lj := &lumberjack.Logger{
		Filename:   filepath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	mw := io.MultiWriter(os.Stderr, lj)
	serviceLogger := logger.NewServiceLogger(mw, "app")
	log.SetOutput(logger.NewStdlibAdapter(serviceLogger)) // redirect stdlib logging to us
	log.SetFlags(0)
	return &serviceLogger
}

// Mechanical stuff
func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}
