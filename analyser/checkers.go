package analyser

import (
	"github.com/LuckeLucky/demo-analyser-csgo/utils"
)

const (
	MAX_ROUNDS_REGULAR = 30
	WIN_ROUNDS_REGULAR = MAX_ROUNDS_REGULAR/2 + 1
)

func (analyser *Analyser) checkValidRoundStartMoney() bool {
	// if the money value is not set, no need to check
	if !analyser.isMoneySet {
		return true
	}

	// between 0 - 30 rounds start money is 800
	if analyser.roundsPlayed < 30 {
		return analyser.currentStartMoney == 800
	}

	return false
}

func (analyser *Analyser) checkMatchHalf() bool {
	if analyser.roundsPlayed == MAX_ROUNDS_REGULAR/2 {
		return true
	}
	return false
}

func (analyser *Analyser) checkMatchEnd() bool {
	ctScore, tScore := analyser.ctScore, analyser.tScore
	roundsInOvertime := ctScore + tScore - MAX_ROUNDS_REGULAR

	if ((ctScore == WIN_ROUNDS_REGULAR) != (tScore == WIN_ROUNDS_REGULAR)) || roundsInOvertime >= 0 {
		absDiff := utils.Abs(ctScore - tScore)
		x := roundsInOvertime % analyser.overtimeMaxRounds
		nRoundsOfOTHalf := analyser.overtimeMaxRounds / 2
		if roundsInOvertime < 0 || ((x == 0 && absDiff == 2) || (x > nRoundsOfOTHalf && absDiff >= nRoundsOfOTHalf)) {
			return true
		}
	}

	return false
}
