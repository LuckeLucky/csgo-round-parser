package analyser

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleParse(t *testing.T) {
	//Demo use RoundEndofficial instead of roundend
	f, err := os.Open("../test-demos/mlp_demo_round_end_official.dem")
	assert.Equal(t, nil, err)

	a := NewAnalyser(f)
	a.ParseToEnd()

	assert.Equal(t, 22, a.roundsPlayed)
	assert.Equal(t, 2, len(a.halfs))

	assert.Equal(t, 3, a.halfs[0].halfCtScore)
	assert.Equal(t, 12, a.halfs[0].halfTScore)

	assert.Equal(t, 4, a.halfs[1].halfCtScore)
	assert.Equal(t, 3, a.halfs[1].halfTScore)

}
