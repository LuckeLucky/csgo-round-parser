package analyser

import (
	"fmt"

	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

func (analyser *Analyser) handlerKills(e events.Kill) {
	fmt.Print("%v\n", e)
}
