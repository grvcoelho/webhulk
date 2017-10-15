package config

type Server struct {
	ListenOn   string `yaml:"listen_on"`
	HealtCheck string `yaml:"health_check"`
}
