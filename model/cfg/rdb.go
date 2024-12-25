package cfg

type Rdb struct {
	Host                     string `mapstructure:"HOST"`
	Port                     int32  `mapstructure:"PORT"`
	Username                 string `mapstructure:"USERNAME"`
	Password                 string `mapstructure:"PASSWORD"`
	DbName                   string `mapstructure:"DB_NAME"`
	Migration                bool   `mapstructure:"MIGRATION"`
	MaxIdleConns             int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns             int    `mapstructure:"MAX_OPEN_CONNS"`
	ConnectTimeoutInSeconds  int    `mapstructure:"CONNECT_TIMEOUT_IN_SECONDS"`
	ConnMaxLifetimeInMinutes int    `mapstructure:"CONN_MAX_LIFETIME_IN_MINUTES"`
}
