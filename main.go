package main

import (
	"log/slog"
	"os"
	"runtime"

	"github.com/thebravebyte/findr/db"
	"github.com/thebravebyte/findr/httpserver"
	"github.com/thebravebyte/findr/internals"
	"github.com/thebravebyte/findr/internals/controller"
	"github.com/thebravebyte/findr/pkg/mail"
)

var bufferSize int

func main() {
	// Setting up the logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting the LinkerApp service")

	// mail channel to use to receive data via a channel
	MailChan := make(chan internals.Mail, bufferSize)

	worker := runtime.NumCPU()

	URI := os.Getenv("DB_URI")

	if URI == "" {
		logger.Error("DB_URI is not set or unavailable")
		os.Exit(1)
	}

	logger.Info("Starting the LinkerApp service.....")

	// MailSystem: starting the Goroutine to listen or be active to receive data for the mailing system
	// and have it process at any instance in time
	go mail.MailDelivery(MailChan, worker)
	logger.Info("Linkr Mailing System Starting ........")



	// OpenConnection: function to open a connection to the database
	DBclient := db.OpenConnection("mongodb+srv://findr-admin:findr-admin@cluster0.cygk5u4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", logger)

	// Setting up the server configuration with actually values or details
	srvCfg := &httpserver.SrvConfig{
		Address:        ":8080",
		ReadTimeout:    15,
		WriteTimeout:   15,
		IdleTimeout:    60,
		MaxHeaderBytes: 1 << 20,
	}

	mainServer := httpserver.ConfigServer(srvCfg, logger)

	mainServer.Use(mainServer.AddMiddleware)
	mainServer.Use(mainServer.LoggingAppRequest(logger))

	// import the funtion fro the controller that returns the service interface
	fa := controller.NewFindr(DBclient, logger, MailChan)

	// Call the handlers from routes which havr the services interface passed as parameters
	mainServer.AddRoutes(fa)

	//  Todo: Ask fpr the path for these path of the statics files
	mainServer.StaticFile("/favicon.ico", "./static/favicon.ico")
	mainServer.StaticFile("/static", "./static")
	// Todo: Reminder

	// Starting the application
	mainServer.StartApp()

}
