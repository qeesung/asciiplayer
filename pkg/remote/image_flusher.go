package remote

type ImageFlusherHandler struct {
	BaseFlushHandler
	Filename   string
	ImageCache string
}
