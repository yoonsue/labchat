package entry

import (
	"log"
	"os"
)

func setLog(logpath string) (file *os.File) {
	println("LogFile: " + logpath)
	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(file)
	return file
}

func cleanlog(logpath string) {
	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
}
