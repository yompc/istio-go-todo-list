package dao

import (
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)
var (
 DB *gorm.DB
)
func InitMySQL()(err error) {
	viper.BindEnv("gorm.uri","GORM_URI")
	viper.SetDefault("gorm.dialect", "mysql")
	viper.SetDefault("gorm.uri", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	var dialect = viper.GetString("gorm.dialect")
	var uri = viper.GetString("gorm.uri")

	DB, err = gorm.Open(dialect, uri)
	if err != nil {
		log.Error().Err(err).Caller()
	}
	err = DB.DB().Ping()
	if err != nil {
		log.Error().Err(err).Caller()
	}
	return
}

func Close()  {
	_ = DB.Close()
}
