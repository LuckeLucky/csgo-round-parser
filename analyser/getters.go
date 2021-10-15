package analyser

func (analyser *Analyser) getGameTick() (int, bool) {
	var err bool
	tick := analyser.parser.GameState().IngameTick()
	if tick < 0 {
		err = true
	}
	return tick, err
}

func (analyser *Analyser) getAllPlayersRoundStartMoney() (money int) {
	if analyser.players == nil || len(analyser.players) == 0 {
		return 0
	}

	for _, p := range analyser.players {
		money += p.RoundStartMoney()
	}

	return
}
