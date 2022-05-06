package analyser

func (analyser *Analyser) setDefault() {
	analyser.overtimeMaxRounds = 6
	analyser.freeArmor = 0
	analyser.matchEnded = false
}

func (analyser *Analyser) resetHalfScores() {
	analyser.halfCtScore = 0
	analyser.halfTScore = 0
}

func (analyser *Analyser) switchSideScores() {
	analyser.ctScore, analyser.tScore = analyser.tScore, analyser.ctScore
}

func (analyser *Analyser) SetDefaultConvarConfig() {
	convarsConfig := make(map[string]int)
	convarsConfig["regularStartMoney"] = 800
	convarsConfig["overtimeStartMoney"] = 16000

	analyser.convarsConfig = convarsConfig
}
