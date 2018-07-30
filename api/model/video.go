package model

import (
	"database/sql"
	"log"
	"time"
	"github.com/Zereker/video_server/api/utils"
	"github.com/Zereker/video_server/api/conn"
)

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCTime string
}

func AddNewVideo(aid int, name string) (*VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	cTime := t.Format("Jan 02 2006, 15:04:05")
	stmt, err := conn.DBConn.Prepare("insert into video_info(id, author_id, name, display_ctime) values (?,?,?,?)")
	if err != nil {
		log.Printf("AddNewVideo, err: %s", err)
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vid, aid, name, cTime)
	if err != nil {
		log.Printf("AddNewVideo, err: %s", err)
		return nil, err
	}
	res := &VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCTime: cTime}
	return res, nil
}

func GetVideoInfo(vid string) (*VideoInfo, error) {
	stmt, err := conn.DBConn.Prepare("select author_id,name,display_ctime from video_info where id = ? and status = 0")
	if err != nil {
		log.Printf("GetVideoInfo, err: %s", err)
		return nil, err
	}
	defer stmt.Close()

	var aid int
	var name string
	var cTime string
	err = stmt.QueryRow(vid).Scan(&aid, &name, &cTime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetVideoInfo, err: %s", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	res := &VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCTime: cTime}
	return res, nil
}

func DeleteVideo(vid string) error {
	stmt, err := conn.DBConn.Prepare("update video_info set status = 1 where id = ?")
	if err != nil {
		log.Printf("DeleteVideo, err: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vid)
	if err != nil {
		log.Printf("DeleteVideo, err: %s", err)
		return err
	}
	return nil
}
