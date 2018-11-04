package decoder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGifDecoder_DecodeFromFile test decode gif file to multi frames
func TestGifDecoder_DecodeFromFile(t *testing.T) {
	assertions := assert.New(t)
	gifDecoder := NewGifDeCoder()
	imageFilename := "testdata/suolong.gif"
	frames, err := gifDecoder.DecodeFromFile(imageFilename)
	assertions.Nil(err)
	assertions.NotNil(frames)
	assertions.Equal(len(frames), 52)
}
