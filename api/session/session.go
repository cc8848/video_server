package session

import (
	"strconv"
	"log"
	"database/sql"
	"sync"
	"github.com/Zereker/video_server/api/conn"
)

type SimpleSession struct {
	Username string
	TTL      int64
}

func InsertSession(sid string, ttl int64, username string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmt, err := conn.DBConn.Prepare("insert into session (session_id, TTL, username) values (?,?,?)")
	if err != nil {
		log.Printf("InsertSession, err: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(sid, ttlStr, username)
	if err != nil {
		log.Printf("InsertSession, err: %s", err)
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*SimpleSession, error) {
	simpleSession := SimpleSession{}
	stmt, err := conn.DBConn.Prepare("select TTL,username from session where session_id = ?")
	if err != nil {
		log.Printf("RetrieveSession, err: %s", err)
		return nil, err
	}
	defer stmt.Close()
	var ttl string
	var username string
	err = stmt.QueryRow(sid).Scan(&ttl, &username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err != nil {
		return nil, err
	} else {
		simpleSession.TTL = res
		simpleSession.Username = username
	}
	return &simpleSession, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmt, err := conn.DBConn.Prepare("SELECT * FROM session")
	if err != nil {
		log.Printf("RetrieveAllSessions, err %s", err)
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("RetrieveAllSessions, err %s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlStr string
		var username string
		if err := rows.Scan(&id, &ttlStr, &username); err != nil {
			log.Printf("RetrieveAllSessions, err %s", err)
			break
		}

		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
			ss := &SimpleSession{Username: username, TTL: ttl}
			m.Store(id, ss)
			log.Printf(" session id: %s, ttl: %d", id, ss.TTL)
		}

	}

	return m, nil
}

func DeleteSession(sid string) error {
	stmt, err := conn.DBConn.Prepare("DELETE FROM session WHERE session_id = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmt.Query(sid); err != nil {
		return err
	}

	return nil
}
