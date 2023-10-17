package apiserver

type Config struct {
	BindAddr      string `toml:"bind_addr"`
	LogLevel      string `toml:"log_level"`
	StanUrl       string `toml:"stan_url"`
	StanClusterId string `toml:"stan_clusted_id"`
	StanClientId  string `toml:"stan_client_id"`
	DatabaseURL   string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:      ":8080",
		LogLevel:      "debug",
		StanUrl:       "nats://localhost:1234",
		StanClusterId: "test-cluster",
		StanClientId:  "test-client",
		DatabaseURL:   "host=localhost user=sletkov password=postgres dbname=ordersDB sslmode=disable",
	}
}
