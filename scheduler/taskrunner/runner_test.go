package taskrunner

import (
	"testing"
	"log"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %d\n", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Execute received: %d\n", d)
			default:
				break forloop
			}
		}
		return nil
	}

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(5 * time.Second)
}
