package settings

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	//"tests":         "../../settings/tests.json",
}

type Configuration struct {
	Database     string       `yaml:"database,omitempty"`
	BaseUrl      string       `yaml:"base_url,omitempty"`
	BasePort     string       `yaml:"base_port,omitempty"`
	LogFile      string       `yaml:"logfile,omitempty"`
	Debug        bool         `yaml:"debug"`
	RedisAddress string       `yaml:"redis_address,omitempty"`
	RedisAuth    string       `yaml:"redis_auth,omitempty"`
	JWTSettings  Settings  `yaml:"jwt_settings,omitempty"`
}

type Settings struct {
	PrivateKeyPath     string `yaml:"private_key_path"`
	PublicKeyPath      string `yaml:"public_key_path"`
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "preproduction"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file")
	}
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}
