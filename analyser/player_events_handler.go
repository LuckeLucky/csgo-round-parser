package analyser

import (
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func (analyser *Analyser) handlerKills(e events.Kill) {
	// Maybe need to handle roundend official
	if !analyser.inRound {
		return
	}

	killer, victim, assister := e.Killer, e.Victim, e.Assister

	isHS := e.IsHeadshot

	if killer == nil || victim == nil {
		return
	}

	analyser.players[killer.SteamID64].AddKill()
	analyser.players[victim.SteamID64].AddDeath()

	if isHS {
		analyser.players[killer.SteamID64].AddHeadShot()
	}

	if assister == nil {
		return
	}

	analyser.players[assister.SteamID64].AddAssist()

}
