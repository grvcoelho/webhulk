package config

type Server struct {
	ListenOn    string `yaml:"listen_on"`
	HealthCheck string `yaml:"health_check"`
}
