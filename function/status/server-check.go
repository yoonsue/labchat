package status

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	model "github.com/yoonsue/labchat/model/status"
)

// Service declares the methods that status service provides.
type Service interface {
	ServerCheck() model.Server
}

type service struct {
	statusServer model.Server
}

// NewService return struct which provides Service interface
func NewService(statusServer model.Server) Service {
	return &service{
		statusServer: statusServer,
	}
}

// ServerCheck returns server status like temperature and request time
func (s *service) ServerCheck() model.Server {
	log.Println("end newserver")
	s.statusServer.Temperature = getTemp()
	log.Println("start getUptime function")
	s.statusServer.Uptime = getUptime(s.statusServer.BootTime)
	return s.statusServer
}

// // Got information here :https://www.kernel.org/doc/Documentation/thermal/sysfs-api.txt
// const temperatureFile = "/sys/class/thermal/thermal_zone0/temp"

// // getTemp returns server temperature
// func getTemp() model.Temperature {
// 	data, err := ioutil.ReadFile(temperatureFile)
// 	if err != nil {
// 		log.Println(errors.Wrap(err, "failed to read temperature file"))
// 		return -1
// 	}

// 	// TO BE IMPLEMENTED : HOW TO DEAL WITH 'linebreak'
// 	fileAsString := string(data)
// 	fileLines := strings.Split(fileAsString, "\n")
// 	fileAsInt, _ := strconv.Atoi(fileLines[0])

// 	// Unit: millidegree Celsius
// 	temp := model.Temperature(fileAsInt / 1000)

// 	if temp <= 0 {
// 		log.Printf("temperature is lower than zero: %s\n", temp.String())
// 		return -1
// 	}
// 	return temp
// }

// ExampleCmdOutput return output of terminal command
func ExampleCmdOutput(cmd string) ([]byte, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return out, nil
}

// getTemp returns server temperature CentOS
func getTemp() model.Temperature {
	data, err := ExampleCmdOutput("sensors | grep \"Core 0\"")
	if err != nil {
		log.Println(errors.Wrap(err, "failed to execute 'sensors' command"))
		return -1
	}

	// TO BE IMPLEMENTED : HOW TO DEAL WITH 'linebreak'
	byteToString := string(data)
	fileLines := strings.Split(byteToString, "\n")

	re := regexp.MustCompile("[0-9]+")
	linestring := re.FindAllString(fileLines[0], -1)
	tempInt, _ := strconv.Atoi(linestring[1])

	// Unit: millidegree Celsius
	temp := model.Temperature(tempInt)

	if temp <= 0 {
		log.Printf("temperature is lower than zero: %s\n", temp.String())
		return -1
	}
	return temp
}

func getUptime(bootTime time.Time) time.Duration {
	since := time.Since(bootTime)
	log.Println("in getUptime function")
	return since
}
