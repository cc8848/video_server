package model

import "testing"

var tempVideoId string

func TestVideoWorkFlow(t *testing.T) {
	clearTable()
	t.Run("PrepareUser", testAddUserCredential)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempVideoId = video.Id
}

func testGetVideoInfo(t *testing.T) {
	videoInfo, err := GetVideoInfo(tempVideoId)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
	if videoInfo.AuthorId != 1 {
		t.Errorf("Expect get %d, but get %d", 1, videoInfo.AuthorId)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempVideoId)
	if err != nil {
		t.Errorf("Error of DeleteVideo: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempVideoId)
	if err != nil || video != nil {
		t.Errorf("Error of ReGetVideoInfo: %v", err)
	}
}
