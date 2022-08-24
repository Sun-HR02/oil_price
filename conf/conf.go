package conf

import (
	"github.com/spf13/viper"
	"log"
)

func SetConf() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("conf")   // path to look for the config file in
}
func DBconf() (string, string, string) {
	SetConf()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	dbport := viper.GetString("mysql.dbport")
	dburl := viper.GetString("mysql.dburl")
	dbpasswd := viper.GetString("mysql.dbpasswd")

	return dburl, dbport, dbpasswd
}

func AppPort() string {
	SetConf()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	port := viper.GetString("app.port")
	return port
}
