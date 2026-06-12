package config

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
