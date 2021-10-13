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
				analyser.currentStartMoney, _ = strconv.ParseFloat(cvar.Value, 64)
				analyser.isMoneySet = true
			} else if cvar.Name == "mp_free_armor" {
				analyser.freeArmor, _ = strconv.Atoi(cvar.Value)
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
