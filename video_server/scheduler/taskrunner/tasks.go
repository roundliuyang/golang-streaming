package taskrunner

import (
	"errors"
	"github.com/alanhou/golang-streaming/video_server/scheduler/dbops"
	"github.com/alanhou/golang-streaming/video_server/scheduler/ossops"
	"log"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) // 简单起见读3条
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func deleteVideo(vid string) error {

	ossfn := "videos/" + vid
	bn := "liuyang-oss"
	ok := ossops.DeleteObject(ossfn, bn)
	if !ok {
		log.Printf("Deleting video error, oss operation failed")
		return errors.New("deleting video error")
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
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
