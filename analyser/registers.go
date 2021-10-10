package analyser

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"

func (analyser *Analyser) registerMatchEventHandlers() {

	//Round End Related Events
	analyser.parser.RegisterEventHandler(func(e events.RoundEnd) { analyser.handleRoundEnd(e) })
	analyser.parser.RegisterEventHandler(func(e events.ScoreUpdated) { analyser.handleRoundEnd(e) })
}
