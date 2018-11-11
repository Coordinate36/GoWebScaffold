package config

import (
	"os"
)

type Configuration struct {
	ServerAddr string

	DbDriver string
	DbUri    string
}

var devConfig = Configuration{

	// base config
	ServerAddr: ":14000",
	DbDriver:   "sqlite3",
	DbUri:      "platform.db",
}

var prodConfig = Configuration{

	// base config
	ServerAddr: ":14000",
	DbDriver:   "sqlite3",
	DbUri:      "platform.db",
}

var Config Configuration

func init() {
	env, ok := os.LookupEnv("PLATFORM_ENV")
	if !ok {
		env = "local"
	}

	Config = map[string]Configuration{
		"local": LocalConfig,
		"dev":   devConfig,
		"prod":  prodConfig,
	}[env]
}
