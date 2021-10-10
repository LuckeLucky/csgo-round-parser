package analyser

import (
	"fmt"

	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func (analyser *Analyser) handleRoundEnd(e interface{}) {
	tick, err := analyser.getGameTick()
	if err { //negative tick
		return
	}

	if !analyser.checkMatchValidity() {
		return
	}

	switch e.(type) {
	case events.ScoreUpdated:
		fmt.Println("ScoreUpdated")
		eventScoreUpdated := e.(events.ScoreUpdated)

		if eventScoreUpdated.TeamState == nil || eventScoreUpdated.TeamState.Opponent == nil {
			return
		}

		switch eventScoreUpdated.TeamState.Team() {
		case common.TeamCounterTerrorists:
			analyser.ctScore++
		case common.TeamTerrorists:
			analyser.tScore++
		}

	case events.RoundEnd:
		fmt.Println("RoundEnd")
		eventRoundEnd := e.(events.RoundEnd)

		if !analyser.checkFinishedRoundValidity(eventRoundEnd) {
			return
		}

		switch eventRoundEnd.WinnerState.Team() {
		case common.TeamCounterTerrorists:
			analyser.ctScore++

		case common.TeamTerrorists:
			analyser.tScore++

		default:
			analyser.isCancelled = true
		}
	} // end switch e.(type)

}
