package common

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Env     string
	Version string
	Port    int

	Db struct {
		Dsn     string
		SslMode string
	}
}

func (c *Config) Parse() {
	flag.StringVar(&c.Env, "env", c.defaultEnv(), "Environment (development|staging|production)")
	flag.StringVar(&c.Version, "version", c.defaultVersion(), "API version")
	flag.IntVar(&c.Port, "port", c.defaultPort(), "API server port")

	flag.StringVar(&c.Db.Dsn, "dsn", c.defaultDsn(), "PostgreSQL DSN")
	flag.StringVar(&c.Db.SslMode, "ssl-mode", c.defaultSslMode(), "PostgreSQL SSL requirement")

	flag.Parse()
}

func (c *Config) Validate() error {
	if c.Db.Dsn == "" {
		return errors.New("the 'dsn' flag is required")
	}

	return nil
}

func (c *Config) defaultEnv() string {
	const defaultEnv = "development"

	if env, exists := os.LookupEnv("ENV"); exists {
		return env
	}
	return defaultEnv
}

func (c *Config) defaultVersion() string {
	const defaultVer = "1.0.0"

	if version, exists := os.LookupEnv("VERSION"); exists {
		return version
	}
	return defaultVer
}

func (c *Config) defaultPort() int {
	const defaultPort = 8080

	if portEnv, exists := os.LookupEnv("PORT"); exists {
		port, err := strconv.Atoi(portEnv)
		if err == nil {
			return port
		}
	}
	return defaultPort
}

func (c *Config) defaultDsn() string {
	const defaultDsn = ""

	if dsn, exists := os.LookupEnv("DSN"); exists {
		return dsn
	}
	return defaultDsn
}

func (c *Config) defaultSslMode() string {
	const defaultSslMode = "disable"

	if sslMode, exists := os.LookupEnv("SSL_MODE"); exists {
		return sslMode
	}
	return defaultSslMode
}
