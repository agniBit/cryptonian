package cfg

type Twilio struct {
	AccountSID  string `mapstructure:"ACCOUNT_SID"`
	AuthToken   string `mapstructure:"AUTH_TOKEN"`
	PhoneNumber string `mapstructure:"PHONE_NUMBER"`
}
