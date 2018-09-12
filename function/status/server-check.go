package status

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	model "github.com/yoonsue/labchat/model/status"
)

// ServerCheck returns server status like temperature and request time
func ServerCheck(t time.Time) *model.Server {
	s := model.NewServer()
	s.Temperature = getTemp()
	s.Uptime = getUptime(t)
	return s
}

// Got information here :https://www.kernel.org/doc/Documentation/thermal/sysfs-api.txt
const temperatureFile = "/sys/class/thermal/thermal_zone0/temp"

// getTemp returns server temperature
func getTemp() model.Temperature {
	data, err := ioutil.ReadFile(temperatureFile)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to read temperature file"))
		return -1
	}

	// TO BE IMPLEMENTED : HOW TO DEAL WITH 'linebreak'
	fileAsString := string(data)
	fileLines := strings.Split(fileAsString, "\n")
	fileAsInt, _ := strconv.Atoi(fileLines[0])

	// Unit: millidegree Celsius
	temp := model.Temperature(fileAsInt / 1000)

	if temp <= 0 {
		log.Printf("temperature is lower than zero: %s\n", temp.String())
		return -1
	}
	return temp
}

func getUptime(modTime time.Time) time.Duration {
	since := time.Since(modTime)
	return since
}
