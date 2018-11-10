package remote

import (
	"errors"
	"fmt"
	"github.com/qeesung/asciiplayer/pkg/decoder"
	"github.com/qeesung/image2ascii/convert"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// ImageFlushHandler extends the BaseFlushHandler and responsible for flush the
// ASCII image to remote client.
type ImageFlushHandler struct {
	BaseFlushHandler
	Filename       string
	ImageCache     string
	convertOptions convert.Options
}

// NewImageFlusherHandler create a new image flusher handler
func NewImageFlusherHandler(filename string, convertOptions *convert.Options) FlushHandler {
	return &ImageFlushHandler{
		Filename:       filename,
		convertOptions: *convertOptions,
	}
}

// Init for ImageFlushHandler init the image flush handler that is responsible for
// decoding the image and convert the image to ASCII image, then cache the ASCII image.
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

// HandlerFunc for image flush handler return a simple handler function that write the cached ASCII image to remote client.
func (handler *ImageFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("User-Agent"), "curl") {
			http.Redirect(w, r, "https://github.com/qeesung/asciiplayer", http.StatusFound)
			return
		}
		fmt.Fprintln(w, handler.ImageCache)
	}
}
