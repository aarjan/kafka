package config

// AppConfig stores application wide config
type AppConfig struct {
	// App
	ESHost        string `envconfg:"es_host"`
	ESPort        string `envconfg:"es_port"`
	ListenHost    string `envconfig:"server_host" `
	ListenPort    uint32 `envconfig:"server_port" `
	Debug         bool   `envconfig:"server_debug"`
	LogFormatJSON bool   `envconfig:"log_format_json"`

	Index   string   `envconfig:"es_index" required:"true"`
	Brokers []string `envconfig:"kafka_brokers" required:"true"`
}
