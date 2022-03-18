package utils

import (
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func PadSpaceEnd(text string, count int) string {
	if len(text) >= count {
		return text
	}
	return text + strings.Repeat(" ", count-len(text))
}

func ReadConfigFile() {
	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath("../")    // set the path of your config file
	viper.AddConfigPath(".")      // set the path of your config file

	err := viper.ReadInConfig()
	CheckError(err)
}
