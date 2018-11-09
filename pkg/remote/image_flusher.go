package remote

import (
	"errors"
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ImageFlushHandler struct {
	BaseFlushHandler
	Filename       string
	ImageCache     string
	convertOptions convert.Options
}

func NewImageFlusherHandler(filename string, convertOptions *convert.Options) FlushHandler {
	return &ImageFlushHandler{
		Filename:       filename,
		convertOptions: *convertOptions,
	}
}

func (handler *ImageFlushHandler) Init() error {
	logrus.Debug("Init the image flush handler...")
	logrus.Debugf("Building the image decoder by %s...", handler.Filename)
	imageDecoder, supported := decoder.NewDecoder(handler.Filename)
	if !supported {
		return errors.New("not supported file type")
	}

	logrus.Debugf("Decoding the image %s...", handler.Filename)
	frames, err := imageDecoder.DecodeFromFile(handler.Filename, nil)
	if err != nil {
		return err
	}

	if len(frames) != 1 {
		return errors.New("extract too many frames from image")
	}

	convertOptions := handler.convertOptions
	converter := convert.NewImageConverter()

	logrus.Debugf("Converting the image %s...", handler.Filename)
	handler.ImageCache = converter.Image2ASCIIString(frames[0], &convertOptions)
	logrus.Debug("Init the image flush handler successfully!")
	return nil
}

func (handler *ImageFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, handler.ImageCache)
	}
}
