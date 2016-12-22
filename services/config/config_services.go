package config

import (
	"os"
	"github.com/auth-web-tokens/settings"
)

var Config settings.Configuration
var LogFile *os.File