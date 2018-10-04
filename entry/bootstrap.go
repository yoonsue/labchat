package entry

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/mgo.v2"

	"github.com/pkg/errors"
	birthdayFunction "github.com/yoonsue/labchat/function/birthday"
	menuFunction "github.com/yoonsue/labchat/function/menu"
	phoneFunction "github.com/yoonsue/labchat/function/phone"
	statusFunction "github.com/yoonsue/labchat/function/status"
	birthdayModel "github.com/yoonsue/labchat/model/birthday"
	menuModel "github.com/yoonsue/labchat/model/menu"
	phoneModel "github.com/yoonsue/labchat/model/phone"
	statusModel "github.com/yoonsue/labchat/model/status"
	"github.com/yoonsue/labchat/repository/inmem"
	"github.com/yoonsue/labchat/repository/mongo"
	"github.com/yoonsue/labchat/server"
)

// defaultConfigPath is the default location where labchat looks for
// a configuration file.
// TODO: allow to change configuration file path by command-line interface.
const defaultConfigPath = "./labchat.conf.yaml"
const defaultLogPath = "./labchat.log"
const defaultPhonePath = "./phone.txt"
const defaultBirthdayPath = "./birthday.txt"

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
		menus        menuModel.Repository
		phonebook    phoneModel.Repository
		birthdayList birthdayModel.Repository
	)

	if yamlConfig.Database == "inmem" {
		log.Println("DB: in-memory")
		menus = inmem.NewMenuRepository()
		phonebook = inmem.NewPhoneRepository()
		birthdayList = inmem.NewBirthdayRepository()
	} else if yamlConfig.Database == "mongo" {
		log.Println("DB: MongoDB")
		session, err := mgo.Dial(yamlConfig.DBURL)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to establish MongoDB session"))
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		menus, _ = mongo.NewMenuRepository(session)
		phonebook, _ = mongo.NewPhoneRepository(session)
		// birthdayList, _ = mongo.NewBirthdayRepository(session)
		log.Println("create the mongoDB session")
	} else {
		log.Fatalf("unsupported database type: %s", yamlConfig.Database)
	}

	var ms menuFunction.Service
	ms = menuFunction.NewService(menus)
	var ps phoneFunction.Service
	ps = phoneFunction.NewService(phonebook)
	statusServer := statusModel.NewServer()
	var ss statusFunction.Service
	ss = statusFunction.NewService(statusServer)
	var bs birthdayFunction.Service
	bs = birthdayFunction.NewService(birthdayList)

	serverConfig := server.DefaultConfig()
	serverConfig.Address = yamlConfig.Address
	log.Println("make server configuration")

	labchat, err := server.NewServer(serverConfig, ms, ps, ss, bs)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create labchat server"))
	}
	log.Println("create the labchat server")

	// TO BE CONSIDERED: it would be inside of NewServer.
	ps.IntialStore(defaultPhonePath)
	bs.IntialStore(defaultBirthdayPath)

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
