package model

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)

func TestCommentsWorkFlow(t *testing.T) {
	clearTable()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddComments", testAddNewComment)
	t.Run("ListComments", testListComments)
}

func testAddNewComment(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComment(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComment: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	duration, _ := time.ParseDuration("-1h")
	from, _ := strconv.Atoi(strconv.FormatInt(time.Now().Add(duration).Unix(), 10))
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().Unix(), 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for k, v := range res {
		fmt.Printf("comment: %d, %v \n", k, v)
	}
}
