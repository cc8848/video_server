package model

import (
	"testing"
	"github.com/Zereker/video_server/api/conn"
)

func TestMain(m *testing.M) {
	clearTable()
	m.Run()
	clearTable()
}

func clearTable() {
	conn.DBConn.Exec("truncate user")
	conn.DBConn.Exec("truncate video_info")
	conn.DBConn.Exec("truncate comment")
}
