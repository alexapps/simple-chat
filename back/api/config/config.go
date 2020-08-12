package config

// Configset the set of fields fiiled from external confgiguration file
type Configset struct {
	Config struct {
		Version string
	}
	Log struct {
		Level string
	}
	HTTP struct {
		Port string
	}
}

// InitConfiguration -
func InitConfiguration() (conf Configset, err error) {
	// TODO read config set from TOML file

	// TODO remove  hard code
	conf = Configset{
		Config: struct {
			Version string
		}{
			Version: "1",
		},
		Log: struct {
			Level string
		}{
			Level: "INFO",
		},
		HTTP: struct {
			Port string
		}{
			Port: "7073",
		},
	}
	return
}
