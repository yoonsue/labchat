package menu

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Msg request & response
type Msg struct {
	Request  string
	Response string
}

// saveMsg: json file write : for adding new request-response
// loadMsg: json file read : matching with client request and replying the response

// saveJSON json file write : for adding new request-response
func saveJSON(filepath string, req string, res string) {
	// TODO: need to check the list already had
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open json file"))
	}
	defer file.Close()

	n := 0
	msg := Msg{req, res}

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "filed to marshal msg to json"))
	}

	jsonString := string(jsonBytes)

	n, err = file.Write([]byte(jsonString))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to write json file"))
	}

	log.Println(n, " byte saved in ", filepath)

}

//loadJSON json file read : matching with client request and replying the response
func loadJSON(filepath string) bool {
	file, err := os.OpenFile(filepath, os.O_RDONLY|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Println(errors.Wrap(err, "failed to open json file"))
		saveJSON(filepath, "hi", "hello") // create initial file
	}
	defer file.Close()

	n := 0

	fi, err := file.Stat()
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get status json file"))
	}

	var data = make([]byte, fi.Size())

	file.Seek(0, os.SEEK_SET)

	n, err = file.Read(data)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open configuration file"))
	}

	log.Println(n, " byte read from ", filepath)
	log.Println(string(data))
	return true
}
