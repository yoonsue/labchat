package entry

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/mgo.v2"

	"github.com/pkg/errors"
	menuFunction "github.com/yoonsue/labchat/function/menu"
	menuModel "github.com/yoonsue/labchat/model/menu"
	"github.com/yoonsue/labchat/repository/inmem"
	"github.com/yoonsue/labchat/repository/mongo"
	"github.com/yoonsue/labchat/server"
)

// defaultConfigPath is the default location where labchat looks for
// a configuration file.
// TODO: allow to change configuration file path by command-line interface.
const defaultConfigPath = "./labchat.conf.yaml"
const defaultLogPath = "./labchat.log"

// Bootstrap is the entry point for running the labchat server.
// It generates the necessary configuration files and creates the components
// of the system, and injects the dependencies according to its hierarchy.
func Bootstrap() {
	// TODO: load the configuration.
	resource, _ := setLog(defaultLogPath)
	log.Println("bootstrap the labchat service")

	yamlConfig, err := readConfig(defaultConfigPath)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to read configuration file"))
	}
	log.Println("read configuration file")

	var (
		menus menuModel.Repository
	)

	if yamlConfig.Database == "inmem" {
		menus = inmem.NewMenuRepository()
	} else if yamlConfig.Database == "mongo" {
		session, err := mgo.Dial(yamlConfig.DBURL)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to establish MongoDB session"))
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		menus, _ = mongo.NewMenuRepository("mongo", session)
		log.Println("create the mongoDB session")
	} else {
		log.Fatalf("unsupported database type: %s", yamlConfig.Database)
	}

	var ms menuFunction.Service
	ms = menuFunction.NewService(menus)

	serverConfig := server.DefaultConfig()
	serverConfig.Address = yamlConfig.Address
	log.Println("make server configuration")

	labchat, err := server.NewServer(serverConfig, ms)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create labchat server"))
	}
	log.Println("create the labchat server")

	labchat.Start()
	log.Printf("run the labchat server at %s", serverConfig.Address)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	for {
		select {
		case <-sigc:
			log.Println("received stop signal from OS")
			log.Println("good bye :)")
			resource.cleanup()
			return
		}
	}
}

// cleanup is called just before the process terminates normally. The cleanup
// code ensures that the program is terminated gracefully, and the system
// components are shut down in proper order. It save some contexts which can
// be reused the next time of booting labchat.
func (r *Resource) cleanup() {
	// TODO: implementation.
	r.cleanLog()
}
