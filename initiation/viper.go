package initization

import (
	"github.com/spf13/viper"
	"simon.com/recipe/global"
)

func InitConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		global.G_LOG.Errorf("read in config error. err %v", err)
		panic(err)
	}

	err = v.Unmarshal(&global.G_CONFIG)
	if err != nil {
		global.G_LOG.Errorf("unmarshal in config error. err %v", err)
		panic(err)
	}
	return v
}
