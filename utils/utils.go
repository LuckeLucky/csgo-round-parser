package utils

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintDebug(text string) {
	test := true
	if test {
		fmt.Print(text + "\n")
	}
}

func PrintScores(ctName string, tName string, ctScore int, tScore int) {
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("%s vs %s  %d : %d\n", blue(ctName), red(tName), ctScore, tScore)
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
