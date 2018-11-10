// remote package define the operations that how to flush the ASCII image
// to remote client, it would be different flush handler for different picture
// or video type.
package remote

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/player"
	"github.com/qeesung/asciiplayer/pkg/util"
	"github.com/qeesung/image2ascii/convert"
	"net/http"
	"time"
)

// FlushHandler define the basic oprations that flush image to remote server
type FlushHandler interface {
	Init() error
	HandlerFunc() func(w http.ResponseWriter, r *http.Request)
}

// supportedFlushHandlerMatchers register the supported flush handler
// and if the Match function is return true, just call the constructor  
// function to build the flusher handler.
var supportedFlushHandlerMatchers = []struct {
	Match       func(string) bool
	Constructor func(string, *convert.Options) FlushHandler
}{
	{
		Match:       util.IsGif,
		Constructor: NewGifFlushHandler,
	},
	{
		Match:       util.IsSupportedImage,
		Constructor: NewImageFlusherHandler,
	},
}

// NewFlushHandler is factory method to create flush handler
func NewFlushHandler(filename string, options *convert.Options) (handler FlushHandler, supported bool) {
	for _, matcher := range supportedFlushHandlerMatchers {
		if matcher.Match(filename) {
			return matcher.Constructor(filename, options), true
		}
	}
	return nil, false
}

// BaseFlushHandler is a basic flush handler that define some basic opration
type BaseFlushHandler struct {
}

// Init doing nothing in base flush handler
func (handler *BaseFlushHandler) Init() error {
	return nil
}

// HandlerFunc return a empty hadnler function
func (handler *BaseFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

// Flush flush the string to remote client immediately
func (handler *BaseFlushHandler) Flush(w http.ResponseWriter, s string) error {
	fmt.Fprintf(w, s)
	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Fprintf(w, player.ClearScreen)

	// flush to the remote immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
		return nil
	} else {
		return errors.New("can not flush to invalid writer")
	}
}
