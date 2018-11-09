package remote

import (
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"net/http"
)

func NewGifFlushHandler(filename string, convertOptions *convert.Options) FlushHandler {
	return &GifFlushHandler{
		Filename:       filename,
		FrameCache:     make([]string, 0),
		convertOptions: *convertOptions,
	}
}

type GifFlushHandler struct {
	BaseFlushHandler
	Filename       string
	FrameCache     []string
	convertOptions convert.Options
}

func (handler *GifFlushHandler) Init() error {
	logrus.Debug("Init the gif flush handler...")
	gifDecoder, supported := decoder.NewDecoder(handler.Filename)
	if !supported {
		return errors.New("not supported file type")
	}

	frames, err := gifDecoder.DecodeFromFile(handler.Filename)
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

func (handler *GifFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for {
			for _, frameStr := range handler.FrameCache {
				handler.Flush(w, frameStr)
			}
		}
	}
}
