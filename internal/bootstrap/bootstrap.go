package bootstrap

import (
	"fmt"
)

var appConfig = config{}

func init() {
	setEnvironment()

	fmt.Println(appConfig)
}

func Execute() {
	fmt.Println(appConfig)
}
