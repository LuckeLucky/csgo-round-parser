package analyser

import (
	p_common "github.com/LuckeLucky/demo-analyser-csgo/common"
)

func (analyser *Analyser) setPlayers() {
	for _, p := range analyser.parser.GameState().Participants().Playing() {
		analyser.players[p.SteamID64] = &p_common.Player{Player: p}
	}
}

func (analyser *Analyser) setNewHalf() {
	analyser.halfs = append(analyser.halfs, &Half{
		ctName:      analyser.parser.GameState().TeamCounterTerrorists().ClanName(),
		tName:       analyser.parser.GameState().TeamTerrorists().ClanName(),
		halfCtScore: analyser.halfCtScore,
		halfTScore:  analyser.halfTScore,
	})
}

func (analyser *Analyser) setRound(tick int) {
	analyser.roundStarted = false
	analyser.currentRound.endTick = tick
	analyser.rounds = append(analyser.rounds, analyser.currentRound)
	analyser.roundsPlayed++

	analyser.currentRound = nil

}
