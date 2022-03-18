package analyser

import (
	"log"
	"os"
	"testing"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	utils.ReadConfigFile()
	//Demo use RoundEndofficial instead of roundend
	f, err := os.Open("../test-demos/mlp_demo_round_end_official.dem")
	assert.Equal(t, nil, err)

	a := NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 22, a.roundsPlayed)
	assert.Equal(t, 2, len(a.halfs))

	assert.Equal(t, 3, a.halfs[0].halfCtScore)
	assert.Equal(t, 12, a.halfs[0].halfTScore)

	assert.Equal(t, 4, a.halfs[1].halfCtScore)
	assert.Equal(t, 3, a.halfs[1].halfTScore)

}

func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
		log.SetOutput(os.Stderr)
	}
}
