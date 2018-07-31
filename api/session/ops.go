package session

import (
	"sync"
	"log"
	"time"
	"github.com/Zereker/video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sessionId string) {
	sessionMap.Delete(sessionId)
	err := DeleteSession(sessionId)
	if err != nil {
		log.Printf("deleteExpiredSession, err: %s", err)
	}
}
func LoadSessionsFromDB() {
	sessions, err := RetrieveAllSessions()
	if err != nil {
		log.Printf("LoadSessionsFromDB, err: %s", err)
		return
	}
	sessions.Range(func(key, value interface{}) bool {
		ss := value.(*SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(username string) string {
	uuid, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000
	simpleSession := &SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(uuid, simpleSession)
	err := InsertSession(uuid, ttl, username)
	if err != nil {
		log.Printf("GenerateNewSessionId, err: %s", err)
	}
	return uuid
}

func IsSessionExpired(sessionId string) (string, bool) {
	ss, ok := sessionMap.Load(sessionId)
	if ok {
		ct := nowInMilli()
		if ss.(*SimpleSession).TTL < ct {
			deleteExpiredSession(sessionId)
			return "", true
		}
		return ss.(*SimpleSession).Username, false
	}
	return "", true
}
