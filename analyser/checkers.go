package analyser

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"

func (analyser *Analyser) checkMatchValidity() bool {
	_, err := analyser.getGameTick()

	if analyser.hasMatchStarted || !analyser.hasMatchEnded || err {
		return false
	}
	return true
}

func (analyser *Analyser) checkFinishedRoundValidity(e events.RoundEnd) bool {
	reason := e.Reason
	if reason == events.RoundEndReasonCTSurrender || reason == events.RoundEndReasonDraw {
		return false
	}
	return true
}
