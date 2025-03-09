package cfg

type Binance struct {
	APIKey    string `mapstructure:"API_KEY"`
	APISecret string `mapstructure:"API_SECRET"`
	WsURL     string `mapstructure:"WS_URL"`
}
