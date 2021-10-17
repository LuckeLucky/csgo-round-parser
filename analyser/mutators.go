package analyser

func (analyser *Analyser) setDefault() {
	analyser.overtimeMaxRounds = 6
	analyser.freeArmor = 0
}

func (analyser *Analyser) resetHalfScores() {
	analyser.halfCtScore = 0
	analyser.halfTScore = 0
}
