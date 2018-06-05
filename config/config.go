package config

// AppConfig stores application wide config
type AppConfig struct {
	// App
	ListenPort    uint32   `envconfig:"server_port" required:"true"`
	ListenHost    string   `envconfig:"server_host" required:"true"`
	Debug         bool     `envconfig:"server_debug"`
	LogFormatJSON bool     `envconfig:"log_format_json"`
	Brokers       []string `envconfig:"brokers" required:"true"`
	Index         string   `envconfig:"index" required:"true"`
}
