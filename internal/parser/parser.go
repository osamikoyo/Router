package parser

import "github.com/BurntSushi/toml"

type Line struct {
	LocalhostPort uint `toml:"localhost_port"`
	Hostname string `toml:"hostname"`
}

type Config struct {
	Port uint `toml:"proxy_port"`
	Logging bool `toml:"logging"`
	Lines []Line `toml:"line"`
}

type Parser struct {
	ConfigFilePath string
}

func New() Parser {
	return Parser{
		ConfigFilePath: "config.toml",
	}
}

func (p Parser) Parse() (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(p.ConfigFilePath, &cfg)
	return cfg, err
}