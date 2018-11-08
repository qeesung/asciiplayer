package remote

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/qeesung/asciiplayer/pkg/player"
	"net/http"
	"time"
)

type FlushHandler interface {
	Init() error
	HandlerFunc() func(w http.ResponseWriter, r *http.Request)
}

type BaseFlushHandler struct {
}

func (handler *BaseFlushHandler) Init() error {
	return nil
}

func (handler *BaseFlushHandler) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

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
