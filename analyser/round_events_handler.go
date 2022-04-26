package analyser

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type Half struct {
	ctName string
	tName  string

	halfCtScore int
	halfTScore  int
}

type Round struct {
	startTick       int
	endTick         int
	endOfficialTick int
}

type LiquipediaPlayer struct {
	Page          string
	SteamID       uint64
	CrosshairCode string
}

func (analyser *Analyser) handlerRoundStart(e interface{}) {
	tick, err := analyser.getGameTick()
	if err {
		return
	}

	if analyser.checkMatchEnded() {
		return
	}

	// Rounds Time Limit equal to 1m55s == 115s
	switch switchEvents := e.(type) {
	case events.RoundStart:
		if switchEvents.TimeLimit != 115 {
			return
		}
	case events.MatchStartedChanged:
		if !switchEvents.NewIsStarted {
			return
		}
	case events.RoundFreezetimeEnd:
		if analyser.roundStarted {
			return
		}
	}

	if !analyser.checkValidRoundStartMoney() {
		return
	}
	if !analyser.checkFreeArmor() {
		return
	}
	if !analyser.checkFirstRoundStartEquipmentValue() {
		return
	}
	analyser.roundStarted = true
	analyser.currentRound = &Round{startTick: tick}

}

func (analyser *Analyser) handlerRoundEnd(e events.RoundEnd) {
	tick, err := analyser.getGameTick()
	if err {
		return
	}

	if !analyser.roundStarted {
		return
	}

	if analyser.roundsPlayed == 0 {
		printTableHeader(analyser.parser.GameState().TeamTerrorists().ClanName() + "_vs_" + analyser.parser.GameState().TeamCounterTerrorists().ClanName() + "__" + analyser.mapName)
		printHeaders()
		for _, p := range analyser.parser.GameState().Participants().Playing() {
			analyser.liquipediaPlayers = append(analyser.liquipediaPlayers, LiquipediaPlayer{
				Page:          p.Name,
				SteamID:       p.SteamID64,
				CrosshairCode: p.CrosshairCode(),
			})
			printRow(p.Name, p.SteamID64, p.CrosshairCode())
		}
		printEndTable()
	}

	switch e.Winner {
	case common.TeamCounterTerrorists:
		analyser.halfCtScore++
		analyser.ctScore = e.WinnerState.Score() + 1
		analyser.tScore = e.LoserState.Score()
	case common.TeamTerrorists:
		analyser.halfTScore++
		analyser.tScore = e.WinnerState.Score() + 1
		analyser.ctScore = e.LoserState.Score()
	}

	analyser.printScore()
	analyser.setRoundEnd(tick)
	analyser.checkForMatchHalfOrEnd()
}

func (analyser *Analyser) handlerRoundEndOfficial(e events.RoundEndOfficial) {
	tick, err := analyser.getGameTick()
	if err {
		return
	}

	//Round started and RoundEnd was not dispatched
	if analyser.roundStarted && analyser.currentRound.endTick == 0 {
		//RondEndOfficial is only dispatched after RoundEnd
		//at this point if RoundEnd was dispatched RondEndOfficial will not be processed because roundStarted is false
		//Ct won the round
		if analyser.parser.GameState().TeamCounterTerrorists().Score() > analyser.ctScore {
			analyser.halfCtScore++
			analyser.ctScore = analyser.parser.GameState().TeamCounterTerrorists().Score()
			analyser.tScore = analyser.parser.GameState().TeamTerrorists().Score()
			//t won the round
		} else if analyser.parser.GameState().TeamTerrorists().Score() > analyser.tScore {
			analyser.halfTScore++
			analyser.tScore = analyser.parser.GameState().TeamTerrorists().Score()
			analyser.ctScore = analyser.parser.GameState().TeamCounterTerrorists().Score()
		}
		analyser.printScore()
		analyser.setRoundEndOfficial(tick)
		analyser.checkForMatchHalfOrEnd()
		return
	}

	//RoundEnd was dispatched so we just store the tick of RoundEndOfficial
	if !analyser.roundStarted && analyser.previousRound.endTick != 0 && analyser.previousRound.endOfficialTick == 0 {
		analyser.previousRound.endOfficialTick = tick
	}
}

func (analyser *Analyser) handlerSideSwitch() {
	//Switch our registed sideScores
	analyser.switchSideScores()
}
