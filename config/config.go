package config

type DatabaseConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Database string
}

type Config struct {
	Db DatabaseConfig
}

func ProvideConfig() Config {
	return Config{
		Db: DatabaseConfig{
			User:     "operator",
			Pass:     "pass",
			Host:     "localhost",
			Port:     "3306",
			Database: "my_system",
		},
	}
}
