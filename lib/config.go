package gitbackup

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Config represents the configuration file that will be used to find backup
// targets and where to back up the repositories
type Config struct {
	BackupDirectory string              `yaml:"backup_directory"`
	Targets         []map[string]string `yaml:"targets"`
}

// GetConfig retrieves the configuration file specified by the user and parses
// it into a Config structure.
func GetConfig(configPath string) Config {
	// Make sure the config file exists.
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("The config file %s doesn't exist: %s", configPath, err)
	}

	// Read the config.
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("The config file cannot be read: %s", err)
	}

	// Parse the config.
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("The config file cannot be parsed: %s", err)
	}

	return config
}
