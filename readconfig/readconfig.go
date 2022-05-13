package readconfig

import (
	"strings"

	"github.com/spf13/viper"
)

func ReadConfig(fileKey string) interface{} {
	index := strings.Index(fileKey, ".")
	fileName := fileKey[0:index]
	key := fileKey[index+1:]

	viper := viper.New()
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	return viper.Get(key)
}
