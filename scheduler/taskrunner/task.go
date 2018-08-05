package taskrunner

import (
	"github.com/Zereker/video_server/scheduler/dbops"
	"log"
	"errors"
	"os"
	"sync"
)

func deleteVideo(vid string) error {
	err := os.Remove(VideoPath + vid)
	if err != nil && !os.IsNotExist(err) {
		log.Printf("deleteVideo error: %v\n", err)
		return err
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("VideoClearDispatcher, error: %v\n", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("all tasks finished")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					log.Printf("VideoClearDispatcher, error: %v\n", err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}
