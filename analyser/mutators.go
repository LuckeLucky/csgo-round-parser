package analyser

import (
	"bufio"

	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
)

func (analyser *Analyser) setDefault() {
	analyser.overtimeMaxRounds = 6
	analyser.freeArmor = 0
}

func (analyser *Analyser) resetParser() {
	newStream := bufio.NewReader(analyser.buf)
	parser := demoinfocs.NewParserWithConfig(newStream, analyser.cfg)
	analyser.parser = parser
}

func (analyser *Analyser) resetHalfScores() {
	analyser.halfCtScore = 0
	analyser.halfTScore = 0
}
