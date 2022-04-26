package analyser

import (
	"fmt"
	"strconv"
)

func printTableHeader(title string) {
	fmt.Println("{| class='wikitable'\n!colspan='3'|" + title)
}

func printHeaders() {
	fmt.Println("|-\n!Name!!SteamID!!Crosshair")
}

func printRow(name string, steamId uint64, crosshairCode string) {
	fmt.Println("|-\n|[[" + name + "]]||" + strconv.FormatUint(steamId, 10) + "||" + crosshairCode)
}

func printEndTable() {
	fmt.Println("|}")
}
