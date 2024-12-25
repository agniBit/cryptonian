package cfg

type Auth struct {
	AccessSecret     string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret    string `mapstructure:"REFRESH_SECRET"`
	AccessTTLMinutes int64  `mapstructure:"ACCESS_TTL"`
	RefreshTTLHour   int64  `mapstructure:"REFRESH_TTL"`
}

func (a *Auth) GetAccessSecret() string {
	return a.AccessSecret
}

func (a *Auth) GetRefreshSecret() string {
	return a.RefreshSecret
}
