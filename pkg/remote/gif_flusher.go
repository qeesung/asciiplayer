package remote

import (
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// GifFlushHandler extends from BaseFlushHandler, and responsible for flushing the gif
// frames to remote client.
type GifFlushHandler struct {
	BaseFlushHandler
	Filename       string
	FrameCache     []string
	convertOptions convert.Options
}

// NewGifFlushHandler create a new gif flush handler
func NewGifFlushHandler(filename string, convertOptions *convert.Options) FlushHandler {
	return &GifFlushHandler{
		Filename:       filename,
		FrameCache:     make([]string, 0),
		convertOptions: *convertOptions,
	}
}

// Init for gif flush handler responsible for decoding the gif to frames
// then decoding the frames to ASCII string slices, then cache results, reduce resource consumption.
func (handler *GifFlushHandler) Init() error {
	logrus.Debug("Init the gif flush handler...")
	gifDecoder, supported := decoder.NewDecoder(handler.Filename)
	if !supported {
		return errors.New("not supported file type")
	}

	frames, err := gifDecoder.DecodeFromFile(handler.Filename, nil)
	if err != nil {
		return err
	}

	convertOptions := handler.convertOptions
	converter := convert.NewImageConverter()

	for _, frame := range frames {
		frameStr := converter.Image2ASCIIString(frame, &convertOptions)
		handler.FrameCache = append(handler.FrameCache, frameStr)
	}
	return nil
}

// HandlerFunc for gif flush handler flush the cached ASCII string slices slice by slice at a centian frequency
func (handler *GifFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("User-Agent"), "curl") {
			http.Redirect(w, r, "https://github.com/qeesung/asciiplayer", http.StatusFound)
			return
		}
		for {
			for _, frameStr := range handler.FrameCache {
				handler.Flush(w, frameStr)
			}
		}
	}
}
