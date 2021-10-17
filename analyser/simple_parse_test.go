package analyser

import (
	"log"
	"os"
	"testing"

	"github.com/LuckeLucky/demo-analyser-csgo/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	utils.ReadConfigFile()
	f, err := os.Open("../test-demos/Caretos Gaming_vs_Rhyno Esports_de_nuke.dem")
	assert.Equal(t, nil, err)

	a := NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 20, a.roundsPlayed)
	assert.Equal(t, 2, len(a.halfs))

	assert.Equal(t, 4, a.halfs[0].halfCtScore)
	assert.Equal(t, 11, a.halfs[0].halfTScore)

	assert.Equal(t, 5, a.halfs[1].halfCtScore)
	assert.Equal(t, 0, a.halfs[1].halfTScore)

	// knife round
	f, err = os.Open("../test-demos/SAW_vs_Galaxy Racer_de_dust2.dem")
	assert.Equal(t, nil, err)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 22, a.roundsPlayed)
	assert.Equal(t, 2, len(a.halfs))

	assert.Equal(t, 13, a.halfs[0].halfCtScore)
	assert.Equal(t, 2, a.halfs[0].halfTScore)

	assert.Equal(t, 4, a.halfs[1].halfCtScore)
	assert.Equal(t, 3, a.halfs[1].halfTScore)

	//knife round
	f, err = os.Open("../test-demos/SAW_vs_DBL PONEY_de_overpass.dem")
	assert.Equal(t, nil, err)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 25, a.roundsPlayed)
	assert.Equal(t, 2, len(a.halfs))

	assert.Equal(t, 6, a.halfs[0].halfCtScore)
	assert.Equal(t, 9, a.halfs[0].halfTScore)

	assert.Equal(t, 7, a.halfs[1].halfCtScore)
	assert.Equal(t, 3, a.halfs[1].halfTScore)

	//no start money set
	f, err = os.Open("../test-demos/SAW_vs_Galaxy Racer_de_inferno.dem")
	assert.Equal(t, nil, err)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 36, a.roundsPlayed)
	assert.Equal(t, 4, len(a.halfs))

	assert.Equal(t, 5, a.halfs[0].halfCtScore)
	assert.Equal(t, 10, a.halfs[0].halfTScore)

	assert.Equal(t, 5, a.halfs[1].halfCtScore)
	assert.Equal(t, 10, a.halfs[1].halfTScore)

	assert.Equal(t, 2, a.halfs[2].halfCtScore)
	assert.Equal(t, 1, a.halfs[2].halfTScore)

	assert.Equal(t, 1, a.halfs[3].halfCtScore)
	assert.Equal(t, 2, a.halfs[3].halfTScore)

	// overtime 16000
	f, err = os.Open("../test-demos/unique-vs-nexus-m3-dust2.dem")
	assert.Equal(t, nil, err)

	viper.Set("overtimeStartMoney", 16000)
	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 36, a.roundsPlayed)
	assert.Equal(t, 4, len(a.halfs))

	assert.Equal(t, 7, a.halfs[0].halfCtScore)
	assert.Equal(t, 8, a.halfs[0].halfTScore)

	assert.Equal(t, 7, a.halfs[1].halfCtScore)
	assert.Equal(t, 8, a.halfs[1].halfTScore)

	assert.Equal(t, 2, a.halfs[2].halfCtScore)
	assert.Equal(t, 1, a.halfs[2].halfTScore)

	assert.Equal(t, 3, a.halfs[3].halfCtScore)
	assert.Equal(t, 0, a.halfs[3].halfTScore)

	assert.Equal(t, 800, a.currentStartMoney)
	assert.Equal(t, 16000, a.currentOvertimeStartMoney)

	// overtime 10000 and missing round start
	f, err = os.Open("../test-demos/SpiritOfAmiga_vs_EshMartinsen_de_mirage.dem")
	assert.Equal(t, nil, err)

	viper.Set("overtimeStartMoney", 10000)

	a = nil
	a = NewAnalyser(f)
	a.SimpleRun()

	assert.Equal(t, 48, a.roundsPlayed)
	assert.Equal(t, 8, len(a.halfs))

	assert.Equal(t, 6, a.halfs[0].halfCtScore)
	assert.Equal(t, 9, a.halfs[0].halfTScore)

	assert.Equal(t, 6, a.halfs[1].halfCtScore)
	assert.Equal(t, 9, a.halfs[1].halfTScore)

	assert.Equal(t, 2, a.halfs[2].halfCtScore)
	assert.Equal(t, 1, a.halfs[2].halfTScore)

	assert.Equal(t, 2, a.halfs[3].halfCtScore)
	assert.Equal(t, 1, a.halfs[3].halfTScore)

	assert.Equal(t, 2, a.halfs[4].halfCtScore)
	assert.Equal(t, 1, a.halfs[4].halfTScore)

	assert.Equal(t, 2, a.halfs[5].halfCtScore)
	assert.Equal(t, 1, a.halfs[5].halfTScore)

	assert.Equal(t, 1, a.halfs[6].halfCtScore)
	assert.Equal(t, 2, a.halfs[6].halfTScore)

	assert.Equal(t, 2, a.halfs[7].halfCtScore)
	assert.Equal(t, 1, a.halfs[7].halfTScore)

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
