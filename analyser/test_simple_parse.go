package analyser

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	defer quiet()
	f, err := os.Open("./test-demos/Caretos Gaming_vs_Rhyno Esports_de_nuke.dem")
	assert.Equal(t, err, nil)

	a1 := NewAnalyser(f)
	a1.SimpleRun()

	assert.Equal(t, a1.roundsPlayed, 20)
	assert.Equal(t, len(a1.halfs), 2)

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
