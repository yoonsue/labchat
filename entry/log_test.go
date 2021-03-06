package entry

import (
	"os"
	"testing"
)

// I thought that I don't have to deal with testing the fatal error
// cause it incidently exit its current situation.

func TestSetLog(t *testing.T) {
	tmpLog := "../tmpLog.log"
	setLog(tmpLog)
	if _, err := os.Stat(tmpLog); os.IsNotExist(err) {
		t.Fatal("tmpLog does not exist")
	}
	defer os.Remove(tmpLog)
}

func TestCleanLog(t *testing.T) {
	r, _ := setLog("../tmpLog.log")
	r.cleanLog()
	if _, err := os.Stat(r.logpath); os.IsNotExist(err) {
		t.Fatal("tmpLog does not exist")
	}
	defer os.Remove(r.logpath)
}
