package config

import (
	"os"
	"strconv"
	"time"

	"github.com/yousefggg/common-lib/pkg/errors"
)

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
	AuthConfig     AuthConfig
}

type AppConfig struct {
	Port        string
	Environment string
	LogLevel    string
}

type DatabaseConfig struct {
	URL          string
	MaxOpenConns int //Макс БД подключенно 
	MaxIdleConns int //Максимум БД в ожидании
	ConnTimeout  time.Duration
}

type AuthConfig struct {
	JWTSecret string
	TokenTime time.Duration
}
//CFG_MISSING - нет переменной 
//CFG_INVALID - Ошибка конвертации 

func LoadConfig() (*Config, error) {
	cfg := &Config{
		AppConfig:      AppConfig{},
		DatabaseConfig: DatabaseConfig{},
		AuthConfig:     AuthConfig{},
	}

	cfg.AppConfig.Port = os.Getenv("APP_PORT")
	if cfg.AppConfig.Port == "" {
		return nil, errors.NewErr("CFG_MISSING", "APP_PORT is required", nil)
	}
	cfg.AppConfig.Environment = os.Getenv("APP_ENVIRONMENT")
	cfg.AppConfig.LogLevel = os.Getenv("APP_LOG_LEVEL")

	cfg.DatabaseConfig.URL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseConfig.URL == "" {
		return nil, errors.NewErr("CFG_MISSING", "DATABASE_URL is required", nil)
	}

	valMaxOpen := os.Getenv("DATABASE_MAX_OPEN_CONNS")
	if valMaxOpen == "" {
		return nil, errors.NewErr("CFG_MISSING", "DATABASE_MAX_OPEN_CONNS is required", nil)
	}
	maxOpen, err := strconv.Atoi(valMaxOpen)
	if err != nil {
		return nil, errors.NewErr("CFG_INVALID", "DATABASE_MAX_OPEN_CONNS must be integer", err)
	}
	cfg.DatabaseConfig.MaxOpenConns = maxOpen

	valMaxIdle := os.Getenv("DATABASE_MAX_IDLE_CONNS")
	if valMaxIdle == "" {
		return nil, errors.NewErr("CFG_MISSING", "DATABASE_MAX_IDLE_CONNS is required", nil)
	}
	maxIdle, err := strconv.Atoi(valMaxIdle)
	if err != nil {
		return nil, errors.NewErr("CFG_INVALID", "DATABASE_MAX_IDLE_CONNS must be integer", err)
	}
	cfg.DatabaseConfig.MaxIdleConns = maxIdle

	valTimeout := os.Getenv("DATABASE_CONN_TIMEOUT")
	if valTimeout == "" {
		return nil, errors.NewErr("CFG_MISSING", "DATABASE_CONN_TIMEOUT is required", nil)
	}
	timeout, err := time.ParseDuration(valTimeout)
	if err != nil {
		return nil, errors.NewErr("CFG_INVALID", "DATABASE_CONN_TIMEOUT invalid format", err)
	}
	cfg.DatabaseConfig.ConnTimeout = timeout

	cfg.AuthConfig.JWTSecret = os.Getenv("AUTH_JWT_SECRET")
	if cfg.AuthConfig.JWTSecret == "" {
		return nil, errors.NewErr("CFG_MISSING", "AUTH_JWT_SECRET is required", nil)
	}

	valTokenTime := os.Getenv("AUTH_TOKEN_TIME")
	if valTokenTime == "" {
		return nil, errors.NewErr("CFG_MISSING", "AUTH_TOKEN_TIME is required", nil)
	}
	tokenTime, err := time.ParseDuration(valTokenTime)
	if err != nil {
		return nil, errors.NewErr("CFG_INVALID", "AUTH_TOKEN_TIME invalid format", err)
	}
	cfg.AuthConfig.TokenTime = tokenTime

	return cfg, nil
}