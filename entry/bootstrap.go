package entry

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/server"
)

// Bootstrap is the entry point for running the labchat server.
// It generates the necessary configuration files and creates the components
// of the system, and injects the dependencies according to its hierarchy.
func Bootstrap() {
	// TODO: load the configuration.
	log.Println("bootstrap the labchat service")

	log.Println("create the labchat server")
	serverConfig := server.DefaultConfig()
	labchat, err := server.NewServer(serverConfig)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create labchat server"))
	}

	log.Println("run the labchat server")
	labchat.Start()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	for {
		select {
		case <-sigc:
			log.Println("received stop signal from OS")
			log.Println("good bye :)")
			cleanup()
			return
		}
	}
}

// cleanup is called just before the process terminates normally. The cleanup
// code ensures that the program is terminated gracefully, and the system
// components are shut down in proper order. It save some contexts which can
// be reused the next time of booting labchat.
func cleanup() {
	// TODO: implementation.
}
