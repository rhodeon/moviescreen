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
		Dsn          string
		SslMode      string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
}

func (c *Config) Parse() {
	flag.StringVar(&c.Env, "env", c.defaultEnv(), "Environment (development|staging|production)")
	flag.StringVar(&c.Version, "version", c.defaultVersion(), "API version")
	flag.IntVar(&c.Port, "port", c.defaultPort(), "API server port")

	flag.StringVar(&c.Db.Dsn, "db-dsn", c.defaultDsn(), "PostgreSQL DSN")
	flag.StringVar(&c.Db.SslMode, "db-ssl-mode", c.defaultSslMode(), "PostgreSQL SSL requirement")
	flag.IntVar(&c.Db.MaxOpenConns, "db-max-open-conns", c.defaultMaxOpenConns(), "PostgreSQL maximum number of open connections")
	flag.IntVar(&c.Db.MaxIdleConns, "db-max-idle-conns", c.defaultMaxIdleConns(), "PostgreSQL maximum number of idle connections")
	flag.StringVar(&c.Db.MaxIdleTime, "db-max-idle-time", c.defaultMaxIdleTime(), "PostgreSQL maximumn idle time")

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

func (c *Config) defaultMaxOpenConns() int {
	const defMaxOpenConns = 25

	if maxOpenConns, exists := os.LookupEnv("DB_MAX_OPEN_CONNS"); exists {
		port, err := strconv.Atoi(maxOpenConns)
		if err == nil {
			return port
		}
	}
	return defMaxOpenConns
}

func (c *Config) defaultMaxIdleConns() int {
	const defMaxIdleConns = 25

	if maxIdleConns, exists := os.LookupEnv("DB_MAX_IDLE_CONNS"); exists {
		port, err := strconv.Atoi(maxIdleConns)
		if err == nil {
			return port
		}
	}
	return defMaxIdleConns
}

func (c *Config) defaultMaxIdleTime() string {
	const defMaxIdleTime = "15m"

	if maxIdleTime, exists := os.LookupEnv("DB_MAX_IDLE_TIME"); exists {
		return maxIdleTime
	}
	return defMaxIdleTime
}
