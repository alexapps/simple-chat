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
	conf := Configset{
		Config: Config{"1"},
		Log:    Log{"INFO"},
		HTTP:   HTTP{"7073"},
	}
	return
}
