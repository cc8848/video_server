package taskrunner

const (
	ReadToDispatch = "d"
	ReadToExecute  = "e"
	Close          = "c"

	VideoPath = "./videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
