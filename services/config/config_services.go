package config

import (
	"os"
	"github.com/auth-web-tokens/settings"
	"log"
	"strings"
	"path/filepath"
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

var Config settings.Configuration
var LogFile *os.File

func init() {
	var configFile string

	if configFile = os.Getenv("WEB-TOKENS"); configFile == "" {
		configFile = "config.yaml"
	}

	err := ReadConfig(configFile)
	log.Println("Reading config " + configFile)

	if err != nil {
		panic(err)
	}
}

func ReadConfig(configFileName string) error {
	filename, _ := filepath.Abs(configFileName)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	EnableLogfile(Config.LogFile)

	return err
}

func EnableLogfile(logfileName string) *os.File {
	if logfileName == "" {
		log.Printf("logfile is STDOUT")
		return nil
	}

	log.Printf("logfile is %s", logfileName)
	logFile := logfileName
	logfileNameSlice := strings.Split(logfileName, string(filepath.Separator))

	//try to create log folder
	if len(logfileNameSlice) > 1 {
		logfileNameSlice = logfileNameSlice[:len(logfileNameSlice) - 1]
		logPath := strings.Join(logfileNameSlice, string(filepath.Separator))
		os.Mkdir(logPath, 0777)
	}

	f, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	LogFile = f

	return f
}
