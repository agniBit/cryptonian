package cfg

type ControlSwitch struct {
	KillSwitchEnabled bool          `mapstructure:"KILL_SWITCH_ENABLED"`
	OnlyPaperTrade    bool          `mapstructure:"ONLY_PAPER_TRADE"`
	UpdateOrderInDB   bool          `mapstructure:"UPDATE_ORDER_IN_DB"`
	XP                map[string]XP `mapstructure:"XP"`
}

type XP struct {
	Enabled bool   `mapstructure:"ENABLED"`
	Control string `mapstructure:"CONTROL"`
	Test    string `mapstructure:"TEST"`
	Variant string `mapstructure:"VARIANT"`
}
