package env

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetPort() (port string) {
	port = fmt.Sprintf(":%s", viper.GetString("PORT"))
	return
}
