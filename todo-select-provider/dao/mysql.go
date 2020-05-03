package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func InitMySql()(err error)  {
	viper.BindEnv("gorm.uri","GORM_URI")
	viper.SetDefault("gorm.dialect", "mysql")
	viper.SetDefault("gorm.uri", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	var dialect = viper.GetString("gorm.dialect")
	var uri = viper.GetString("gorm.uri")

	DB, err = gorm.Open(dialect, uri)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return DB.DB().Ping()
}

func Close()  {
	DB.Close()
}
