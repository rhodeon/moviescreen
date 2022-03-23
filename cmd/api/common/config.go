package common

import "flag"

type Config struct {
	Env     string
	Version string
	Port    int
}

func (c *Config) Parse() {
	flag.StringVar(&c.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&c.Version, "version", "1.0.0", "API version")
	flag.IntVar(&c.Port, "port", 8080, "API server port")

	flag.Parse()
}
