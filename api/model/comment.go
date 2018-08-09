package model

import (
	"log"
	"github.com/Zereker/video_server/api/utils"
	"github.com/Zereker/video_server/api/conn"
)

type Comment struct {
	Id         string
	VideoId    string
	AuthorName string
	Content    string
}

type NewComment struct {
	AuthorId int `json:"author_id"`
	Content string `json:"content"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}


func AddNewComment(vid string, aid int, comment string) error {
	cid, err := utils.NewUUID()
	if err != nil {
		log.Printf("AddNewComment, err: %s", err)
		return err
	}
	stmt, err := conn.DBConn.Prepare("insert into comment(id, video_id, author_id, content) values (?,?,?,?)")
	if err != nil {
		log.Printf("AddNewComment, err: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(cid, vid, aid, comment)
	if err != nil {
		log.Printf("AddNewComment, err: %s", err)
		return err
	}
	return nil
}

func ListComments(vid string, from, to int) ([]*Comment, error) {
	stmt, err := conn.DBConn.Prepare(`select comment.id,user.username,comment.content
					from comment inner join user on comment.author_id = user.id
					where comment.video_id = ? 
					and comment.time > from_unixtime(?) 
					and comment.time <= from_unixtime(?)`)
	if err != nil {
		log.Printf("ListComments, err: %s", err)
		return nil, err
	}
	defer stmt.Close()
	var res []*Comment
	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		log.Printf("ListComments, err: %s", err)
		return nil, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &Comment{Id: id, VideoId: vid, AuthorName: name, Content: content}
		res = append(res, c)
	}
	return res, nil
}
