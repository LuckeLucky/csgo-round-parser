package analyser

import (
	"os"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/gogo/protobuf/proto"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

type Analyser struct {
	parser  demoinfocs.Parser
	mapName string

	//-------------------------

	halfs []*Half
	//Rounds ------------------
	roundHandler *RoundHandler
	roundsPlayed int

	//Current ScoreBoard scores
	ctScore int
	tScore  int
	//-------------------------

	//Match flags -------------

	//Convars set
	isMoneySet bool
	//-------------------------

	//Convars -----------------
	currentStartMoney float64
	overtimeMaxRounds int
	freeArmor         int
	//-------------------------

}

func NewAnalyser(f *os.File) *Analyser {

	cfg := demoinfocs.DefaultParserConfig
	cfg.AdditionalNetMessageCreators = map[int]demoinfocs.NetMessageCreator{
		6: func() proto.Message {
			return new(msg.CNETMsg_SetConVar)
		},
	}

	parser := demoinfocs.NewParserWithConfig(f, cfg)

	analyser := &Analyser{}
	analyser.parser = parser

	return analyser

}

func (analyser *Analyser) handleHeader() {

	header, err := analyser.parser.ParseHeader()
	utils.CheckError(err)
	analyser.mapName = header.MapName
}

func (analyser *Analyser) Run() {
	analyser.handleHeader()
	analyser.setDefault()

	analyser.roundHandler = new(RoundHandler)

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()

	var err error
	for ok := true; ok; ok, err = analyser.parser.ParseNextFrame() {
		utils.CheckError(err)
	}

	analyser.printHalfs()
	analyser.printMap()
}
