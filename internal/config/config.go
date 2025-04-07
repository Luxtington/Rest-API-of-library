package config

type Config struct {
	DbConnectionString string
	ServerHost         string
}

func Load() (*Config, error) {
	return &Config{
		DbConnectionString: "postgres://postgres:postgres@localhost/go_test?sslmode=disable",
		ServerHost:         ":8080",
	}, nil
}
