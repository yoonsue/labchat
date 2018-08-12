package entry

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// yamlConfig contains information for running labchat.
// Read from the file: {config_path}/labchat.conf.yaml
type yamlConfig struct {
	Address  string `json:"address"`
	Database string `json:"database"`
	DBURL    string `json:"DBURL"`
}

// readConfig reads configuration from the configuration file.
func readConfig(fpath string) (*yamlConfig, error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	yc := &yamlConfig{}

	err = yaml.Unmarshal(b, yc)
	if err != nil {
		return nil, err
	}

	return yc, nil
}
