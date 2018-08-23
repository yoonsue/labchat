package entry

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ghodss/yaml"
)

func TestReadConfig(t *testing.T) {
	yc := &yamlConfig{
		Address:  "localhost:2300",
		Database: "inmem", // 'inmem' or 'mongo'
		// DBURL:    "localhost",
	}

	testCases := []struct {
		tmpFile string
	}{
		{
			"client.cfg",
		},
		{
			"",
		},
	}

	for _, c := range testCases {
		tmpFile, err := ioutil.TempFile("", c.tmpFile)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		b, err := yaml.Marshal(yc)
		if err != nil {
			t.Fatal(err)
		}

		_, err = tmpFile.Write(b)
		if err != nil {
			t.Fatal(err)
		}

		cfg, err := readConfig(tmpFile.Name())
		if err != nil {
			t.Error(err)
			return
		}
		if cfg == nil {
			t.Logf("no file name %s", c.tmpFile)
		} else if cfg.Address != yc.Address {
			t.Errorf("expected %s, got %s", yc.Address, cfg.Address)
		}
	}
}
