package function

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/pkg/errors"
	"github.com/yoonsue/labchat/model"
)

// ServerCheck returns server status like temperature and request time
func ServerCheck() model.ServerStatus {
	status := model.ServerStatus{}
	// Access somewhere to get status..
	status.Temperature = getTemp()
	status.TimeStamp = getTime()
	return status
}

// Got information here :https://www.kernel.org/doc/Documentation/thermal/sysfs-api.txt
// /sys/class/thermal/thermal_zone*/temp

// getTemp returns server temperature
func getTemp() model.Temperature {
	data, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get temperature"))
		data = nil
	}

	// Unit: millidegree Celsius
	temp := model.Temperature(float32frombytes(data) / 1000)
	if temp <= 0 {
		log.Println("failed to get temperature")
		return -1
	}
	return temp
}

// getTime returns current time
func getTime() time.Time {
	time := time.Now()
	return time
}

// float64frombytes changes bytes to float64
func float32frombytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
