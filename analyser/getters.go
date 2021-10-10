package analyser

func (analyser *Analyser) getGameTick() (int, bool) {
	var err bool
	tick := analyser.parser.GameState().IngameTick()
	if tick < 0 {
		err = true
	}
	return tick, err
}
