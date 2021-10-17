package analyser

import (
	"log"
	"os"
	"testing"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	defer quiet()
	utils.ReadConfigFile()
	f, err := os.Open("../test-demos/Caretos Gaming_vs_Rhyno Esports_de_nuke.dem")
	assert.Equal(t, err, nil)

	a := NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, a.roundsPlayed, 20)
	assert.Equal(t, len(a.halfs), 2)

	assert.Equal(t, a.halfs[0].halfCtScore, 4)
	assert.Equal(t, a.halfs[0].halfTScore, 11)

	assert.Equal(t, a.halfs[1].halfCtScore, 5)
	assert.Equal(t, a.halfs[1].halfTScore, 0)

	f, err = os.Open("../test-demos/SAW_vs_Galaxy Racer_de_dust2.dem")
	assert.Equal(t, err, nil)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, a.roundsPlayed, 22)
	assert.Equal(t, len(a.halfs), 2)

	assert.Equal(t, a.halfs[0].halfCtScore, 13)
	assert.Equal(t, a.halfs[0].halfTScore, 2)

	assert.Equal(t, a.halfs[1].halfCtScore, 4)
	assert.Equal(t, a.halfs[1].halfTScore, 3)

	f, err = os.Open("../test-demos/SAW_vs_DBL PONEY_de_overpass.dem")
	assert.Equal(t, err, nil)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, a.roundsPlayed, 25)
	assert.Equal(t, len(a.halfs), 2)

	assert.Equal(t, a.halfs[0].halfCtScore, 6)
	assert.Equal(t, a.halfs[0].halfTScore, 9)

	assert.Equal(t, a.halfs[1].halfCtScore, 7)
	assert.Equal(t, a.halfs[1].halfTScore, 3)

	f, err = os.Open("../test-demos/SAW_vs_Galaxy Racer_de_inferno.dem")
	assert.Equal(t, err, nil)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, a.roundsPlayed, 36)
	assert.Equal(t, len(a.halfs), 4)

	assert.Equal(t, a.halfs[0].halfCtScore, 5)
	assert.Equal(t, a.halfs[0].halfTScore, 10)

	assert.Equal(t, a.halfs[1].halfCtScore, 5)
	assert.Equal(t, a.halfs[1].halfTScore, 10)

	assert.Equal(t, a.halfs[2].halfCtScore, 2)
	assert.Equal(t, a.halfs[2].halfTScore, 1)

	assert.Equal(t, a.halfs[3].halfCtScore, 1)
	assert.Equal(t, a.halfs[3].halfTScore, 2)
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
