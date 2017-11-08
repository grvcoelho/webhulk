package config

type Database struct {
	Address    string `yaml:"address"`
	Migrations string `yaml:"migrations"`
}
