package parser

import "github.com/BurntSushi/toml"

type Line struct {
	LocalhostPort uint `toml:"localhost_port"`
	Hostname string `toml:"hostname"`
}

type config struct {
	Port uint `toml:"proxy_port"`
	Logging bool `toml:"logging"`
	Lines []Line
}

type Config struct {
	Port uint
	Logging bool
	Lines map[string]uint
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
	var cfg config
	_, err := toml.DecodeFile(p.ConfigFilePath, &cfg)

	var newcfg Config
	for _, c := range cfg.Lines {
		newcfg.Lines[c.Hostname] = c.LocalhostPort
	}

	return newcfg, err
}