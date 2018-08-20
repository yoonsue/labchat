package entry

import (
	"log"
	"os"
)

// Resource stores the path of log file for cleaning up the exact log file.
type Resource struct {
	logpath string
}

func setLog(logpath string) (*Resource, error) {
	println("LogFile: " + logpath)
	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
	return &Resource{logpath}, nil
}

func (r *Resource) cleanLog() {
	log.Printf("closing %s\n", r.logpath)
	file, err := os.OpenFile(r.logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
}
