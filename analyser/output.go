package analyser

import (
	"fmt"
	"strconv"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/fatih/color"
)

func (analyser *Analyser) printHalfs() {
	if len(analyser.halfs) == 0 {
		return
	}

	var firstCtNme string

	for i, half := range analyser.halfs {
		if i == 0 {
			firstCtNme = half.ctName
			fmt.Printf("Half  |%.6s|%.6s\n", half.ctName, half.tName)
		}
		printScoresHalf(half, i+1, half.ctName == firstCtNme)
	}
}

func (analyser *Analyser) printMap() {
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "Map:%s\n", color.YellowString(analyser.mapName))
	} else {
		fmt.Printf("Map:%s\n", color.YellowString(analyser.mapName))
	}
}

func printScoresHalf(half *Half, nHalf int, isCTOnLeft bool) {
	ctScore := utils.PadSpaceEnd(strconv.Itoa(half.halfCtScore), 6)
	tScore := utils.PadSpaceEnd(strconv.Itoa(half.halfTScore), 6)

	halfPrint := utils.PadSpaceEnd(strconv.Itoa(nHalf), 6)
	format := []string{color.RedString(tScore), color.BlueString(ctScore)}
	if isCTOnLeft {
		format[0], format[1] = format[1], format[0]
	}
	fmt.Printf("%s|", halfPrint)
	if utils.IsWindows() {
		fmt.Fprintf(color.Output, "%s|%s\n", format[0], format[1])
	} else {
		fmt.Printf("%s|%s\n", format[0], format[1])
	}
}
