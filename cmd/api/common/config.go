package common

import (
	"errors"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Env            string
	Version        string
	Port           int
	DisplayVersion bool

	Db struct {
		Dsn          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}

	Limiter struct {
		Enabled bool
		Rps     float64
		Burst   int
	}

	Smtp struct {
		Host     string
		Port     int
		User     string
		Password string
		Sender   string
	}
}

func (c *Config) Parse() {
	flag.StringVar(&c.Env, "env", c.defaultEnv(), "Environment (development|staging|production)")
	flag.IntVar(&c.Port, "port", c.defaultPort(), "API server port")
	flag.BoolVar(&c.DisplayVersion, "version", false, "Display version and build time")

	flag.StringVar(&c.Db.Dsn, "db-dsn", c.defaultDbDsn(), "PostgreSQL DSN")
	flag.IntVar(&c.Db.MaxOpenConns, "db-max-open-conns", c.defaultDbMaxOpenConns(), "PostgreSQL maximum number of open connections")
	flag.IntVar(&c.Db.MaxIdleConns, "db-max-idle-conns", c.defaultDbMaxIdleConns(), "PostgreSQL maximum number of idle connections")
	flag.StringVar(&c.Db.MaxIdleTime, "db-max-idle-time", c.defaultDbMaxIdleTime(), "PostgreSQL maximumn idle time")

	flag.BoolVar(&c.Limiter.Enabled, "limiter-enabled", c.defaultLimiterEnabled(), "Enable rate limiter")
	flag.Float64Var(&c.Limiter.Rps, "limiter-rps", c.defaultLimiterRps(), "Rate limiter maximum requests per second")
	flag.IntVar(&c.Limiter.Burst, "limiter-burst", c.defaultLimiterBurst(), "Rate limiter maximum burst")

	flag.StringVar(&c.Smtp.Host, "smtp-host", c.defaultSmtpHost(), "SMTP hostname")
	flag.IntVar(&c.Smtp.Port, "smtp-port", c.defaultSmtpPort(), "SMTP port")
	flag.StringVar(&c.Smtp.User, "smtp-user", c.defaultSmtpUser(), "SMTP username")
	flag.StringVar(&c.Smtp.Password, "smtp-pass", c.defaultSmtpPassword(), "SMTP password")
	flag.StringVar(&c.Smtp.Sender, "smtp-sender", c.defaultSmtpSender(), "SMTP sender")

	flag.Parse()
}

func (c *Config) Validate() error {
	if c.Db.Dsn == "" {
		return errors.New("the 'db-dsn' flag is required")
	}

	if c.Smtp.Host == "" {
		return errors.New("the 'smtp-host' flag is required")
	}

	if c.Smtp.User == "" {
		return errors.New("the 'smtp-user' flag is required")
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

func (c *Config) defaultPort() int {
	const defaultPort = 4000

	if portEnv, exists := os.LookupEnv("PORT"); exists {
		port, err := strconv.Atoi(portEnv)
		if err == nil {
			return port
		}
	}
	return defaultPort
}

func (c *Config) defaultDbDsn() string {
	const defaultDsn = ""

	if dsn, exists := os.LookupEnv("DB_DSN"); exists {
		return dsn
	}
	return defaultDsn
}

func (c *Config) defaultDbMaxOpenConns() int {
	const defMaxOpenConns = 25

	if maxOpenConnsEnv, exists := os.LookupEnv("DB_MAX_OPEN_CONNS"); exists {
		maxOpenConns, err := strconv.Atoi(maxOpenConnsEnv)
		if err == nil {
			return maxOpenConns
		}
	}
	return defMaxOpenConns
}

func (c *Config) defaultDbMaxIdleConns() int {
	const defMaxIdleConns = 25

	if maxIdleConnsEnv, exists := os.LookupEnv("DB_MAX_IDLE_CONNS"); exists {
		maxIdleConns, err := strconv.Atoi(maxIdleConnsEnv)
		if err == nil {
			return maxIdleConns
		}
	}
	return defMaxIdleConns
}

func (c *Config) defaultDbMaxIdleTime() string {
	const defMaxIdleTime = "15m"

	if maxIdleTime, exists := os.LookupEnv("DB_MAX_IDLE_TIME"); exists {
		return maxIdleTime
	}
	return defMaxIdleTime
}

func (c *Config) defaultLimiterEnabled() bool {
	const defaultEnabled = true

	if enabledEnv, exists := os.LookupEnv("LIMITER_ENABLED"); exists {
		enabled, err := strconv.ParseBool(enabledEnv)
		if err == nil {
			return enabled
		}
	}
	return defaultEnabled
}

func (c *Config) defaultLimiterRps() float64 {
	const defaultRps = 2

	if rpsEnv, exists := os.LookupEnv("LIMITER_RPS"); exists {
		rps, err := strconv.ParseFloat(rpsEnv, 64)
		if err == nil {
			return rps
		}
	}
	return defaultRps
}

func (c *Config) defaultLimiterBurst() int {
	const defaultBurst = 4

	if burstEnv, exists := os.LookupEnv("LIMITER_BURST"); exists {
		burst, err := strconv.Atoi(burstEnv)
		if err == nil {
			return burst
		}
	}
	return defaultBurst
}

func (c *Config) defaultSmtpHost() string {
	if host, exists := os.LookupEnv("SMTP_HOST"); exists {
		return host
	}
	return ""
}
func (c *Config) defaultSmtpPort() int {
	const defaultPort = 587

	if portEnv, exists := os.LookupEnv("SMTP_PORT"); exists {
		port, err := strconv.Atoi(portEnv)
		if err == nil {
			return port
		}
	}
	return defaultPort
}

func (c *Config) defaultSmtpUser() string {
	if user, exists := os.LookupEnv("SMTP_USER"); exists {
		return user
	}
	return ""
}

func (c *Config) defaultSmtpPassword() string {
	if password, exists := os.LookupEnv("SMTP_PASS"); exists {
		return password
	}
	return ""
}

func (c *Config) defaultSmtpSender() string {
	const defaultSender = "Team Moviescreen <no-reply@moviescreen.net>"
	if sender, exists := os.LookupEnv("SMTP_SENDER"); exists {
		return sender
	}
	return defaultSender
}
