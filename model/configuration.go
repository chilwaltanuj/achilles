package model

import (
	"time"
)

// ApplicationConfiguration represents the application configuration.
type ApplicationConfiguration struct {
	Application   string           `mapstructure:"application"`    // to be unique within organization
	ApplicationID int64            `mapstructure:"application_id"` // to be unique within organization
	HttpPort      int              `mapstructure:"http_port"`
	Log           LogConfig        `mapstructure:"log_config"`
	HTTP          ClientHTTPConfig `mapstructure:"http_config"`
}

// LogConfig represents the log configuration.
type LogConfig struct {
	LogInTerminalOverFile bool `mapstructure:"log_in_terminal_over_file"`
	EnableIndentation     bool `mapstructure:"enable_indentation"`

	LogFilePath   string `mapstructure:"log_file_path"`
	LogMaxSizeMB  int    `mapstructure:"log_max_size_mb"`
	LogMaxBackups int    `mapstructure:"log_max_backups"`
	LogMaxAgeDays int    `mapstructure:"log_max_age_days"`
}

// ClientHTTPConfig represents the HTTP client configuration.
type ClientHTTPConfig struct {
	TimeoutDuration              time.Duration `mapstructure:"timeout_in_ms"`
	HystrixCommand               string        `mapstructure:"hystrix_command"`
	CircuitBreakerActiveTimeInMs int           `mapstructure:"circuit_breaker_active_time_in_ms"`
	MaxConcurrent                int           `mapstructure:"max_concurrent"`
	ErrorThreshold               int           `mapstructure:"error_threshold"`
	RetryMax                     int           `mapstructure:"retry_max"`
	RetryBackoffDuration         time.Duration `mapstructure:"retry_backoff_in_ms"`
	RetryJitterDuration          time.Duration `mapstructure:"retry_jitter_in_ms"`
	RetryMaxWaitDuration         time.Duration `mapstructure:"retry_jitter_in_ms"`
	RequestVolumeThreshold       int           `mapstructure:"request_to_volume_threshold"`
}
