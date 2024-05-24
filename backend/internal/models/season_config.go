package models

type SeasonConfig struct {
	Leagues  []League `yaml:"leagues"`
	IceTimes []string `yaml:"ice_time"`
}
