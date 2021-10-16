package analyser

import (
	"bytes"
	"fmt"
	"io"

	p_common "github.com/LuckeLucky/demo-analyser-csgo/common"
	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/gogo/protobuf/proto"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

type Analyser struct {
	parser demoinfocs.Parser

	//buf and cfg will allow for a secon parse
	buf     *bytes.Buffer
	cfg     demoinfocs.ParserConfig
	mapName string

	rounds       []*Round
	currentRound *Round
	roundsPlayed int
	halfs        []*Half
	roundStarted bool

	//Current ScoreBoard scores
	ctScore int
	tScore  int
	//half scores
	halfCtScore int
	halfTScore  int

	//Convars -----------------
	isMoneySet                bool
	currentStartMoney         int
	currentOvertimeStartMoney int
	overtimeMaxRounds         int
	freeArmor                 int

	//Vars for a 2nd parse
	inRound bool
	players []*p_common.Player
}

func NewAnalyser(demostream io.Reader) *Analyser {
	analyser := &Analyser{}
	analyser.buf = &bytes.Buffer{}
	demostream = io.TeeReader(demostream, analyser.buf)
	analyser.cfg = demoinfocs.DefaultParserConfig
	analyser.cfg.AdditionalNetMessageCreators = map[int]demoinfocs.NetMessageCreator{
		6: func() proto.Message {
			return new(msg.CNETMsg_SetConVar)
		},
	}

	parser := demoinfocs.NewParserWithConfig(demostream, analyser.cfg)
	analyser.parser = parser

	return analyser

}

func (analyser *Analyser) handleHeader() {

	header, err := analyser.parser.ParseHeader()
	utils.CheckError(err)
	analyser.mapName = header.MapName
}

func (analyser *Analyser) SimpleRun() {
	analyser.handleHeader()
	analyser.setDefault()

	analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()

	var err error
	for ok := true; ok; ok, err = analyser.parser.ParseNextFrame() {
		utils.CheckError(err)
	}

	analyser.printHalfs()
	analyser.printMap()
	fmt.Printf("Rounds played:%d\n", analyser.roundsPlayed)
}

func (analyser *Analyser) RunAndAnalyse() {
	analyser.resetParser()
	analyser.roundsPlayed = 0

	analyser.registerAnalyseEventHandlers()

	var err error
	for ok := true; ok; ok, err = analyser.parser.ParseNextFrame() {
		utils.CheckError(err)
	}

	fmt.Printf("rounds:%d\n", analyser.roundsPlayed)

}

func (analyser *Analyser) GetDemoNameWithDetails() (name string) {
	name = fmt.Sprintf("%s_vs_%s_%s", analyser.halfs[0].ctName, analyser.halfs[0].tName, analyser.mapName)
	return name
}
