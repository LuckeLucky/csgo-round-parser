package common

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

type Player struct {
	*common.Player
}

func (p *Player) RoundStartMoney() int {
	return p.Money() + p.MoneySpentThisRound()
}
