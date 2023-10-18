package model

// ApplicationConfiguration represents the application configuration.
type ApplicationConfiguration struct {
	Application   string    `mapstructure:"application"`
	ApplicationID int       `mapstructure:"application_id"`
	HttpPort      int       `mapstructure:"http_port"`
	Log           LogConfig `mapstructure:"log"`
}

// LogConfig represents the log configuration.
type LogConfig struct {
	RedirectLogFileToTerminal bool `mapstructure:"redirect_log_File_to_terminal"`

	LogFilePath   string `mapstructure:"log_path"`
	LogMaxSizeMB  int    `mapstructure:"log_max_size_mb"`
	LogMaxBackups int    `mapstructure:"log_max_backups"`
	LogMaxAgeDays int    `mapstructure:"log_max_age_days"`
}