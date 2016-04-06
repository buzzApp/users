package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/forestgiant/portutil"
	"github.com/forestgiant/semver"
	"github.com/gorilla/mux"
)

const (
	//Version represents the semantic version of this service/api
	Version = "0.1.0"
)

func main() {
	// Setup Semantic Version flags
	err := semver.SetVersion(Version)
	if err != nil {
		log.Fatal(err)
	}

	// Check for command line configuration flags
	var (
		logPathUsage = "Path to the service logs."
		logPathPtr   = flag.String("logpath", "", logPathUsage)
	)
	flag.Parse()

	if len(*logPathPtr) == 0 {
		log.Fatal("You must provide a path where log files can be stored.")
	}

	l := getLogger(*logPathPtr)

	// `package log` domain
	l.Info("Initializing app.", "Main")

	//Obtain an available port
	port, err := portutil.GetUniqueTCP()
	if err != nil {
		log.Fatal(err)
	}
	httpAddress := ":" + strconv.Itoa(port)

	// Register service with Stela api
	serviceRegistration(l, port)

	// Mechanical stuff
	errc := make(chan error)
	go func() {
		errc <- interrupt()
	}()

	// Define our app service
	var service UserService
	service = userService{}
	service = userServiceLogginMiddleware{l, service}

	go func() {
		l.Info("Establishing HTTP Bindings", "Main", "addr", httpAddress, "transport", "HTTP/JSON")

		// Create a new mux router
		router := mux.NewRouter()

		const CreateUserPath = "/users"
		router.Handle(CreateUserPath, handleCreateUser(service)).Methods("POST")
		l.Info("New Handler", "Main", "path", CreateUserPath, "type", "POST")

		// register our router and start the server
		http.Handle("/", router)
		errc <- http.ListenAndServe(httpAddress, nil)
	}()

	fmt.Println("Fatal Error", "Main", <-errc)
}
