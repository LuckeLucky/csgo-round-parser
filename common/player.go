package common

import "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"

type Player struct {
	*common.Player

	kills   int
	deaths  int
	assists int

	headShots int
}

func (p *Player) RoundStartMoney() int {
	return p.Money() + p.MoneySpentThisRound()
}

func (p *Player) AddKill() {
	p.kills++
}

func (p *Player) AddDeath() {
	p.deaths++
}

func (p *Player) AddAssist() {
	p.assists++
}

func (p *Player) AddHeadShot() {
	p.headShots++
}

func (p *Player) GetKills() int {
	return p.kills
}

func (p *Player) GetDeaths() int {
	return p.deaths
}

func (p *Player) GetAssists() int {
	return p.deaths
}

func (p *Player) GetHeadShots() int {
	return p.headShots
}
