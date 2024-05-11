package model

import "time"

// ApplicationConfiguration represents the application configuration.
type ApplicationConfiguration struct {
	Application   string            `mapstructure:"application"`    // to be unique within organization
	ApplicationID int64             `mapstructure:"application_id"` // to be unique within organization
	HttpPort      int               `mapstructure:"http_port"`
	Log           LogConfig         `mapstructure:"log_config"`
	HTTP          ClientHTTPConfig  `mapstructure:"http_config"`
	RDBMS         ClientRDBMSConfig `mapstructure:"rdbms_config"`
	Config        ContextConfig     `mapstructure:"context_config"`
}
type ContextConfig struct {
	Timmeout time.Duration `mapstructure:"timeout"`
}

// LogConfig represents the log configuration.
type LogConfig struct {
	LogInTerminalOverFile bool `mapstructure:"log_in_terminal_over_file"`
	EnableIndentation     bool `mapstructure:"enable_indentation"`

	LogLevel      string `mapstructure:"log_level"`
	LogFilePath   string `mapstructure:"log_file_path"`
	LogMaxSizeMB  int    `mapstructure:"log_max_size_mb"`
	LogMaxBackups int    `mapstructure:"log_max_backups"`
	LogMaxAgeDays int    `mapstructure:"log_max_age_days"`
}

// ClientHTTPConfig represents the HTTP client configuration.
type ClientHTTPConfig struct {
	RequestTimeoutDuration     time.Duration `mapstructure:"request_timeout_ms"`
	CircuitBreakerName         string        `mapstructure:"circuit_breaker_name"`
	CircuitBreakerActiveTimeMs int           `mapstructure:"circuit_breaker_active_duration_ms"`
	MaxConcurrentRequests      int           `mapstructure:"max_concurrent_requests"`
	ErrorThresholdPercentage   int           `mapstructure:"error_threshold_percentage"`
	RetryCountMax              int           `mapstructure:"retry_count_max"`
	RetryBackoffDuration       time.Duration `mapstructure:"retry_backoff_ms"`
	RetryJitterDuration        time.Duration `mapstructure:"retry_jitter_ms"`
	RetryMaxWaitDuration       time.Duration `mapstructure:"retry_duration_max"`
	RequestVolumeThreshold     int           `mapstructure:"request_volume_threshold"`
}

type ClientRDBMSConfig struct {
	DSN                        string        `mapstructure:"dsn"`                            // Data Source Name for database connection
	MaxOpenConns               int           `mapstructure:"max_open_conns"`                 // Maximum number of open connections to the database
	MaxIdleConns               int           `mapstructure:"max_idle_conns"`                 // Maximum number of connections in the idle connection pool
	ConnMaxLifetime            time.Duration `mapstructure:"conn_max_lifetime_ms"`           // Maximum amount of time a connection may be reused (expressed in milliseconds)
	CircuitBreakerName         string        `mapstructure:"circuit_breaker_name"`           // Name of the circuit breaker for the database client
	RequestTimeoutDuration     time.Duration `mapstructure:"request_timeout_duration"`       // Request timeout duration for database operations (expressed in milliseconds)
	CircuitBreakerActiveTimeMs int           `mapstructure:"circuit_breaker_active_time_ms"` // Time in milliseconds that the circuit breaker remains active
	MaxConcurrentRequests      int           `mapstructure:"max_concurrent_requests"`        // Maximum number of concurrent requests for the circuit breaker
	ErrorThresholdPercentage   int           `mapstructure:"error_threshold_percentage"`     // Error percentage threshold to trip the circuit breaker
	RequestVolumeThreshold     int           `mapstructure:"request_volume_threshold"`       // Minimum number of requests to consider before tripping the circuit breaker
	SleepWindow                time.Duration `mapstructure:"sleep_window_ms"`                // Duration the circuit breaker stays open before allowing a single test request (expressed in milliseconds)
	RetryCountMax              int           `mapstructure:"retry_count_max"`                // Maximum number of retry attempts
	RetryBackoffMs             time.Duration `mapstructure:"retry_backoff_ms"`               // Initial backoff interval for retries, expressed in milliseconds
}
