package analyser

import (
	"fmt"

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

func (analyser *Analyser) handlerRoundEnd(e interface{}) {
	tick, err := analyser.getGameTick()
	if err {
		return
	}

	if !analyser.roundStarted {
		return
	}

	switch switchEvents := e.(type) {
	case events.RoundEnd:
		switch switchEvents.Winner {
		case common.TeamCounterTerrorists:
			analyser.halfCtScore++
			analyser.ctScore = switchEvents.WinnerState.Score() + 1
			analyser.tScore = switchEvents.LoserState.Score()
		case common.TeamTerrorists:
			analyser.halfTScore++
			analyser.tScore = switchEvents.WinnerState.Score() + 1
			analyser.ctScore = switchEvents.LoserState.Score()
		}
	case events.RoundEndOfficial:
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
	}
	analyser.printScore()
	analyser.setRound(tick)

	isEnd, isHalf := analyser.checkMatchFinished(), analyser.checkMatchHalf()
	if isEnd || isHalf {
		analyser.setNewHalf()
		if isEnd {
			fmt.Println("---Finish---")
			analyser.setMatchEnded()
		} else {
			fmt.Println("---HALF---")
			analyser.resetHalfScores()
		}
	}
}

func (analyser *Analyser) handlerSideSwitch() {
	//Switch our registed sideScores
	analyser.switchSideScores()
}
