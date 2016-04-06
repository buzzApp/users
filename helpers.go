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
	stela "gitlab.fg/go/stela/api"
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

func serviceRegistration(l *logger.ServiceLogger, port int) {
	// Create stela client
	client, err := stela.NewClient(stela.DefaultStelaHTTPAddress)
	if err != nil {
		l.Error("ServiceRegistration", 0, err)
	}

	// Create a service
	service := new(stela.Service)
	service.Name = "app.service.fg"
	service.Port = port

	// Now register with stela
	client.RegisterService(service, func(err error) {
		l.Error("Service Registration", 0, err, nil)
	})

	l.Info("ServiceRegistration", "Stela Client Registered")
}
