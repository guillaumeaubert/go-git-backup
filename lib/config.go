package gitbackup

import(
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BackupDirectory string `yaml:"backup_directory"`
	Targets []map[string]string `yaml:"targets"`
}

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

