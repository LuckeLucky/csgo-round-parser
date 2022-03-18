package analyser

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

func (analyser *Analyser) setMatchEnded() {
	analyser.matchEnded = true
}
