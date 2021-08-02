package boot

import (
	"github.com/BurntSushi/toml"
	"ihtest/library/global"
)

func InitConfig() error {
	_, err := toml.DecodeFile("config/config.toml", &global.Config)
	if err != nil {
		return err
	}

	return nil
}
