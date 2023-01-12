package config

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// EnvEnvironment environment variable to set up application environment
	EnvEnvironment = "ENVIRONMENT"
	// DefEnvProd define production mode
	DefEnvProd = "prod"
	DefEnvDev  = "dev"
	DefEnvTest = "test"
	// EnvDebug environment variable to set up application debug mode, bool: true, false
	EnvDebug = "DEBUG"
	DefDebug = "true"

	// уровень вывода сообщений логгером:
	EnvLoggerLevel = "LOG_LEVEL"
	DefLoggerLevel = "info"

	// Redis
	EnvRedisDbHost = "REDIS_DB_HOST"
	DefRedisDbHost = ""

	EnvRedisDbPort = "REDIS_DB_PORT"
	DefRedisDbPort = "6379"

	EnvRedisDbNumber = "REDIS_DB_NUMBER"
	DefRedisDbNumber = "1"

	EnvRedisDbPassword = "REDIS_DB_PASSWORD"
	DefRedisDbPassword = ""

	EnvTimeout = "CONTEXT_TIMEOUT"
	DefTimeout = "5000"

	EnvMakvesUrl = "MAKVES_URL"
	DefMakvesUrl = ""

	EnvHost    = "HOST"
	DefEnvHost = ""

	EnvPath    = "APP_PATH"
	DefEnvPath = ""

	EnvPort    = "PORT"
	DefEnvPort = "8080"
)

// Config is a main configuration struct for application
type Config struct {
	Environment string
	Debug       bool
	Log         *LogConfig
	Redis       *Redis
	App         *AppConfig
	Timeout     int
	MakvesUrl   string
}

// Redis -- параметры для Redis
type Redis struct {
	Port     string
	Host     string
	Password string
	DbNumber int
}

type LogConfig struct {
	Level string // std: trace, debug, info, warning, error, fatal, panic
}

type AppConfig struct {
	Host           string
	Port           string
	Path           string
	MaxReceiveSize int
	MaxSendSize    int
}

// конфигуратор един для всех! Инстанс, т.к. м.б. нужен в init() пакетов..
var GlCfg *Config

// New returns application config instance
func New() *Config {
	if GlCfg == nil {
		GlCfg = &Config{
			Environment: ToString(LookupEnv(EnvEnvironment, DefEnvDev)),
			Debug:       ToString(LookupEnv(EnvDebug, DefDebug)) == "true",
			Log: &LogConfig{
				Level: ToString(LookupEnv(EnvLoggerLevel, DefLoggerLevel)),
			},
			Redis: &Redis{
				Host:     ToString(LookupEnv(EnvRedisDbHost, DefRedisDbHost)),
				Port:     ToString(LookupEnv(EnvRedisDbPort, DefRedisDbPort)),
				Password: ToString(LookupEnv(EnvRedisDbPassword, DefRedisDbPassword)),
				DbNumber: ToInt(LookupEnv(EnvRedisDbNumber, DefRedisDbNumber)),
			},
			App: &AppConfig{
				Host: ToString(LookupEnv(EnvHost, DefEnvHost)),
				Port: ToString(LookupEnv(EnvPort, DefEnvPort)),
				Path: ToString(LookupEnv(EnvPath, DefEnvPath)),
			},
			Timeout:   ToInt(LookupEnv(EnvTimeout, DefTimeout)),
			MakvesUrl: ToString(LookupEnv(EnvMakvesUrl, DefMakvesUrl)),
		}
	}

	return GlCfg
}

// IsProduction check that application in production mode
func (c *Config) IsProduction() bool {
	return c.Environment == DefEnvProd
}

// IsDevelopment check that application in development mode
func (c *Config) IsDevelopment() bool {
	return c.Environment == DefEnvDev
}

// IsTest check that application in test mode
func (c *Config) IsTest() bool {
	return c.Environment == DefEnvTest
}

func LookupEnv(name string, defVal interface{}) interface{} {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return defVal
}

func ToInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return v
	case string:
		if intVal, err := strconv.Atoi(v); err == nil {
			return intVal
		}
	}
	panic(fmt.Sprintf("Config.New() ERROR! value %v is not INTEGER or not be converted to INT!", val))
}

func ToUInt(val interface{}) uint {
	switch v := val.(type) {
	case uint:
		return v
	case string:
		if intVal, err := strconv.Atoi(v); err == nil {
			return uint(intVal)
		}
	}
	panic(fmt.Sprintf("Config.New() ERROR! value %v is not UINTEGER or not be converted to UINT!", val))
}

func ToInt64(val interface{}) int64 {
	switch v := val.(type) {
	case int64:
		return v
	case string:
		if intVal, err := strconv.Atoi(v); err == nil {
			return int64(intVal)
		}
	}
	panic(fmt.Sprintf("Config.New() ERROR! value %v is not INTEGER or not be converted to INT64!", val))
}

func ToString(val interface{}) string {
	if strVal, ok := val.(string); ok {
		return strVal
	}
	panic(fmt.Sprintf("Config.New() ERROR! value %v is not STRING!", val))
}
