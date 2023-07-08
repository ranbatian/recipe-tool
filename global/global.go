package global

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"simon.com/recipe/config"
)

var (
	G_VIPER  *viper.Viper
	G_DB     *gorm.DB
	G_CONFIG *config.Config
	G_LOG    *logrus.Logger
)
