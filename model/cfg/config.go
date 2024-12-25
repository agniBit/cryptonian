package cfg

type Config struct {
	Auth          *Auth          `mapstructure:"AUTH"`
	Server        *ServerConfig  `mapstructure:"SERVER"`
	Websocket     *Websocket     `mapstructure:"WEBSOCKET"`
	Rdb           *Rdb           `mapstructure:"RDB"`
	Logger        *Logger        `mapstructure:"LOGGER"`
	NewRelic      *NewRelic      `mapstructure:"NEWRELIC"`
	ControlSwitch *ControlSwitch `mapstructure:"CONTROL_SWITCH"`
	Twilio        *Twilio        `mapstructure:"TWILIO"`
	S3            *S3            `mapstructure:"S3"`
	Binance       *Binance       `mapstructure:"BINANCE"`
}