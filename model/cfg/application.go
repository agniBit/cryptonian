package cfg

type NewRelic struct {
	AppName    string `mapstructure:"APP_NAME"`
	LicenseKey string `mapstructure:"LICENSE_KEY"`
}
