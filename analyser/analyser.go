package analyser

import (
	"demo-analyser-csgo/utils/utils"
	"os"

	"github.com/gogo/protobuf/proto"

	demoinfocs "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/msg"
)

type Analyser struct {
	parser demoinfocs.Parser

	mapName string

	//Team Scores -------------
	ctScore int
	tScore  int
	//-------------------------

	//flags -------------------
	hasMatchStarted bool
	hasMatchEnded   bool
	isCancelled     bool
	//-------------------------
}

func NewAnalyser(file *os.File) *Analyser {

	cfg := demoinfocs.DefaultParserConfig
	cfg.AdditionalNetMessageCreators = map[int]demoinfocs.NetMessageCreator{
		6: func() proto.Message {
			return new(msg.CNETMsg_SetConVar)
		},
	}

	parser := demoinfocs.NewParserWithConfig(file, cfg)
	defer parser.Close()

	analyser := &Analyser{parser: parser}

	return analyser

}

func (analyser *Analyser) handleHeader() {
	header, err := analyser.parser.ParseHeader()
	utils.CheckError(err)
	analyser.mapName = header.MapName
}

func (analyser *Analyser) Run() {
	analyser.handleHeader()

	//analyser.registerNetMessageHandlers()
	analyser.registerMatchEventHandlers()
	//analyser.registerFirstPlayerEventHandlers()

	for hasMoreFrames, err := true, error(nil); hasMoreFrames; hasMoreFrames, err = analyser.parser.ParseNextFrame() {
		utils.CheckError(err)
	}
}
