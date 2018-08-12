package taskrunner

const (
	ReadyToDispatch = "d"
	ReadyToExecute  = "e"
	CLOSE           = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
