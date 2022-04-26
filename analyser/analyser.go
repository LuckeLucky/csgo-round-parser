package analyser

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/gogo/protobuf/proto"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

type Match struct {
	Link    string
	Title   string
	Date    string
	Players []LiquipediaPlayer
}
type Analyser struct {
	parser demoinfocs.Parser

	cfg     demoinfocs.ParserConfig
	mapName string

	rounds        []*Round
	currentRound  *Round
	previousRound *Round
	roundsPlayed  int
	halfs         []*Half

	roundStarted bool

	matchEnded bool

	//Current ScoreBoard scores
	ctScore int
	tScore  int
	//half scores
	halfCtScore int
	halfTScore  int

	//Convars -----------------
	isMoneySet                bool
	isOvertimeMoneySet        bool
	currentStartMoney         int
	currentOvertimeStartMoney int
	overtimeMaxRounds         int
	freeArmor                 int

	liquipediaPlayers []LiquipediaPlayer
}

func NewAnalyser(demostream io.Reader) *Analyser {
	analyser := &Analyser{}
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
	analyser.printRoundsPlayed()

	data := Match{
		Players: analyser.liquipediaPlayers,
	}

	file, _ := json.MarshalIndent(data, "", "    ")

	ctName := "cts"
	tName := "ts"
	if len(analyser.halfs) > 0 {
		ctName = analyser.halfs[0].ctName
		tName = analyser.halfs[0].tName
	}

	_ = ioutil.WriteFile("jsons/"+ctName+"_vs_"+tName+"_at_"+analyser.mapName+".json", file, 0644)
}
