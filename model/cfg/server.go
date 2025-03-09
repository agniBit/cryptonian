package cfg

type ServerConfig struct {
	SkipOtpSendError   bool   `mapstructure:"SKIP_OTP_SEND_ERROR"`
	Environment         string `mapstructure:"ENV"`
	Port               string `mapstructure:"PORT"`
	CommonDeadlineInMs int64  `mapstructure:"COMMON_DEADLINE_IN_MS"`
	LogFile            string `mapstructure:"LOG_FILE"`
}
