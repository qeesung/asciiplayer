package remote

import (
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"net/http"
)

func NewGifFlushHandler(filename string) FlushHandler {
	return &GifFlusherHandler{
		Filename:   filename,
		FrameCache: make([]string, 0),
	}
}

type GifFlusherHandler struct {
	BaseFlushHandler
	Filename   string
	FrameCache []string
}

func (handler *GifFlusherHandler) Init() error {
	gifDecoder, supported := decoder.NewDecoder(handler.Filename)
	if !supported {
		return errors.New("not supported file type")
	}

	frames, err := gifDecoder.DecodeFromFile(handler.Filename)
	if err != nil {
		return err
	}

	convertOptions := convert.DefaultOptions
	converter := convert.NewImageConverter()

	for _, frame := range frames {
		frameStr := converter.Image2ASCIIString(frame, &convertOptions)
		handler.FrameCache = append(handler.FrameCache, frameStr)
	}
	return nil
}

func (handler *GifFlusherHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for {
			for _, frameStr := range handler.FrameCache {
				handler.Flush(w, frameStr)
			}
		}
	}
}
