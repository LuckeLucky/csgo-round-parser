package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func CheckError(err error, where ...string) {
	if err != nil {
		if len(where) > 0 {
			fmt.Println("Error detected in: " + where[0])
		}
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
