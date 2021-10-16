package analyser

import (
	"strconv"

	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

func (analyser *Analyser) registerNetMessageHandlers() {
	// Register handler for net messages updates
	analyser.parser.RegisterNetMessageHandler(func(m *msg.CNETMsg_SetConVar) {
		for _, cvar := range m.Convars.Cvars {
			if cvar.Name == "mp_overtime_maxrounds" {
				analyser.overtimeMaxRounds, _ = strconv.Atoi(cvar.Value)
			} else if cvar.Name == "mp_startmoney" {
				analyser.currentStartMoney, _ = strconv.Atoi(cvar.Value)
				analyser.isMoneySet = true
			} else if cvar.Name == "mp_free_armor" {
				analyser.freeArmor, _ = strconv.Atoi(cvar.Value)
			} else if cvar.Name == "mp_overtime_startmoney" {
				/*sometimes mp_overtime_startmoney is used instead of start_money for overtimes*/
				analyser.currentOvertimeStartMoney, _ = strconv.Atoi(cvar.Value)
			}
		}
	})
}

func (analyser *Analyser) registerMatchEventHandlers() {
	//Round start
	analyser.parser.RegisterEventHandler(func(e events.RoundStart) { analyser.handlerRoundStart(e) })
	analyser.parser.RegisterEventHandler(func(e events.MatchStartedChanged) { analyser.handlerRoundStart(e) })

	//Round ends
	analyser.parser.RegisterEventHandler(func(e events.RoundEnd) { analyser.handlerRoundEnd(e) })

}

func (analyser *Analyser) registerRoundEnd(tick int) {
	analyser.roundStarted = false
	analyser.currentRound.endTick = tick
	analyser.rounds = append(analyser.rounds, analyser.currentRound)
	analyser.roundsPlayed++

	analyser.currentRound = nil

}
