package status

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"

	"github.com/pkg/errors"
	model "github.com/yoonsue/labchat/model/status"
)

// ServerCheck returns server status like temperature and request time
func ServerCheck() *model.Server {
	s := model.NewServer()
	s.Temperature = getTemp()
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

	// Unit: millidegree Celsius
	temp := model.Temperature(float32FromBytes(data) / 1000)
	if temp <= 0 {
		log.Printf("temperature is lower than zero: %s\n", temp.String())
		return -1
	}
	return temp
}

// float64frombytes changes bytes to float64
func float32FromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
