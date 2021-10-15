package analyser

import (
	"github.com/LuckeLucky/demo-analyser-csgo/utils"
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
	startTick int
	endTick   int
}

func (analyser *Analyser) handlerRoundStart(e interface{}) {
	tick, err := analyser.getGameTick()
	if err {

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

	switch e.Winner {
	case common.TeamCounterTerrorists:
		utils.PrintScores(e.WinnerState.ClanName(), e.LoserState.ClanName(),
			e.WinnerState.Score()+1, e.LoserState.Score())
		analyser.halfCtScore++
		analyser.ctScore = e.WinnerState.Score() + 1
		analyser.tScore = e.LoserState.Score()
	case common.TeamTerrorists:
		utils.PrintScores(e.LoserState.ClanName(), e.WinnerState.ClanName(),
			e.LoserState.Score(), e.WinnerState.Score()+1)
		analyser.halfTScore++
		analyser.tScore = e.WinnerState.Score() + 1
		analyser.ctScore = e.LoserState.Score()
	}

	analyser.registerRoundEnd(tick)

	isEnd, isHalf := analyser.checkMatchEnd(), analyser.checkMatchHalf()
	if isEnd || isHalf {
		analyser.setNewHalf()
		if isEnd {
			utils.PrintDebug("---Finish---")
		} else {
			utils.PrintDebug("---HALF---")
			analyser.resetHalfScores()
		}
	}
}
