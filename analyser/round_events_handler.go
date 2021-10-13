package analyser

import (
	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type RoundHandler struct {
	roundStarted bool

	halfCtScore int
	halfTScore  int
}

type Half struct {
	ctName string
	tName  string

	halfCtScore int
	halfTScore  int
}

func (analyser *Analyser) handlerRoundStart(e interface{}) {
	_, err := analyser.getGameTick()
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
	analyser.roundHandler.roundStarted = true

}

func (analyser *Analyser) handlerRoundEnd(e events.RoundEnd) {
	_, err := analyser.getGameTick()
	if err {
		return
	}

	if !analyser.roundHandler.roundStarted {
		return
	}

	switch e.Winner {
	case common.TeamCounterTerrorists:
		utils.PrintScores(e.WinnerState.ClanName(), e.LoserState.ClanName(),
			e.WinnerState.Score()+1, e.LoserState.Score())
		analyser.roundHandler.halfCtScore++
		analyser.ctScore = e.WinnerState.Score() + 1
		analyser.tScore = e.LoserState.Score()
	case common.TeamTerrorists:
		utils.PrintScores(e.LoserState.ClanName(), e.WinnerState.ClanName(),
			e.LoserState.Score(), e.WinnerState.Score()+1)
		analyser.roundHandler.halfTScore++
		analyser.tScore = e.WinnerState.Score() + 1
		analyser.ctScore = e.LoserState.Score()
	}

	analyser.roundHandler.roundStarted = false
	analyser.roundsPlayed++

	if analyser.checkMatchEnd() {
		utils.PrintDebug("---Finish---")
		analyser.halfs = append(analyser.halfs, &Half{
			ctName:      analyser.parser.GameState().TeamCounterTerrorists().ClanName(),
			tName:       analyser.parser.GameState().TeamTerrorists().ClanName(),
			halfCtScore: analyser.roundHandler.halfCtScore,
			halfTScore:  analyser.roundHandler.halfTScore,
		})

	} else if analyser.checkMatchHalf() {
		utils.PrintDebug("---HALF---")
		analyser.halfs = append(analyser.halfs, &Half{
			ctName:      analyser.parser.GameState().TeamCounterTerrorists().ClanName(),
			tName:       analyser.parser.GameState().TeamTerrorists().ClanName(),
			halfCtScore: analyser.roundHandler.halfCtScore,
			halfTScore:  analyser.roundHandler.halfTScore,
		})
		analyser.roundHandler.resetHalfScores()
	}

}

func (rh *RoundHandler) resetHalfScores() {
	rh.halfCtScore = 0
	rh.halfTScore = 0
}
