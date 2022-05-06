package analyser

import (
	"fmt"
	"strconv"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/pterm/pterm"
)

func (analyser *Analyser) printHalfs() {
	if len(analyser.halfs) == 0 {
		return
	}

	var firstCtNme string

	for i, half := range analyser.halfs {
		if i == 0 {
			firstCtNme = half.ctName
			fmt.Printf("Half  |%.6s|%.6s\n", utils.PadSpaceEnd(half.ctName, 6), utils.PadSpaceEnd(half.tName, 6))
		}
		printScoresHalf(half, i+1, half.ctName == firstCtNme)
	}
}

func (analyser *Analyser) printMap() {
	pterm.Println("Map: " + pterm.Yellow(analyser.mapName))
}

func (analyser *Analyser) printFinish() {
	fmt.Println("---Finish---")
}

func (analyser *Analyser) printHalf() {
	fmt.Println("---HALF---")
}

func (analyser *Analyser) printRoundsPlayed() {
	fmt.Printf("Rounds played:%d\n", analyser.roundsPlayed)
}

func printScoresHalf(half *Half, nHalf int, isCTOnLeft bool) {
	ctScore := utils.PadSpaceEnd(strconv.Itoa(half.halfCtScore), 6)
	tScore := utils.PadSpaceEnd(strconv.Itoa(half.halfTScore), 6)

	halfPrint := utils.PadSpaceEnd(strconv.Itoa(nHalf), 6)
	format := []string{pterm.Red(tScore), pterm.Blue(ctScore)}
	if isCTOnLeft {
		format[0], format[1] = format[1], format[0]
	}
	fmt.Printf("%s|", halfPrint)
	pterm.Println(format[0] + "|" + format[1])
}
func (analyser *Analyser) printScore() {
	ctName := analyser.parser.GameState().TeamCounterTerrorists().ClanName()
	tName := analyser.parser.GameState().TeamTerrorists().ClanName()
	pterm.Printf("%s vs %s  %d : %d\n", pterm.Blue(ctName), pterm.Red(tName), analyser.ctScore, analyser.tScore)
}
