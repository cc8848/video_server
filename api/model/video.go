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

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
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

func ListVideoInfo(uname string, from, to int) ([]*VideoInfo, error) {
	stmtOut, err := conn.DBConn.Prepare(`SELECT video_info.id, video_info.author_id, video_info.name, video_info.display_ctime FROM video_info 
		INNER JOIN user ON video_info.author_id = user.id
		WHERE user.username = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?) 
		ORDER BY video_info.create_time DESC`)

	var res []*VideoInfo

	if err != nil {
		return res, err
	}

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCTime: ctime}
		res = append(res, vi)
	}

	defer stmtOut.Close()

	return res, nil
}
