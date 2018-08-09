package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/Zereker/video_server/api/model"
	"github.com/Zereker/video_server/api/session"
	"github.com/Zereker/video_server/api/response"
	"github.com/Zereker/video_server/api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	userCredential := &model.User{}
	err := json.Unmarshal(res, userCredential)
	if err != nil {
		log.Printf("CreateUser, err: %s", err)
		response.SendErrorResponse(w, response.RequestBodyParseFailedError)
		return
	}
	err = model.AddUserCredential(userCredential.Username, userCredential.Password)
	if err != nil {
		log.Printf("CreateUser, err: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}
	id := session.GenerateNewSessionId(userCredential.Username)
	su := response.SignedUp{
		Success:   true,
		SessionId: id,
	}
	if resp, err := json.Marshal(su); err != nil {
		log.Printf("CreateUser, err: %s", err)
		response.SendErrorResponse(w, response.InternalFaults)
		return
	} else {
		response.SendNormalResponse(w, string(resp), http.StatusCreated)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	userCredential := &model.User{}
	if err := json.Unmarshal(res, userCredential); err != nil {
		log.Printf("Login, err: %s", err)
		response.SendErrorResponse(w, response.RequestBodyParseFailedError)
		return
	}

	// validate the request body
	username := p.ByName("username")
	if username != userCredential.Username {
		log.Printf("Login url name: %s", username)
		log.Printf("Login body name: %s", userCredential.Username)
		response.SendErrorResponse(w, response.NoAuthUserError)
		return
	}

	password, err := model.GetUserCredential(userCredential.Username)
	if err != nil {
		log.Printf("Login, err: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}

	if len(userCredential.Password) > 0 && password != userCredential.Password {
		log.Printf("Login, err: %s", err)
		response.SendErrorResponse(w, response.NoAuthUserError)
		return
	}

	id := session.GenerateNewSessionId(userCredential.Username)
	su := response.SignedUp{
		Success:   true,
		SessionId: id,
	}
	if resp, err := json.Marshal(su); err != nil {
		log.Printf("Login, err: %s", err)
		response.SendErrorResponse(w, response.InternalFaults)
		return
	} else {
		response.SendNormalResponse(w, string(resp), http.StatusOK)
	}

}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	uname := p.ByName("username")
	u, err := model.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}

	ui := &model.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(ui); err != nil {
		response.SendErrorResponse(w, response.InternalFaults)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}

}

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		log.Printf("Unathorized user \n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &model.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		response.SendErrorResponse(w, response.RequestBodyParseFailedError)
		return
	}

	vi, err := model.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	log.Printf("Author id : %d, name: %s \n", nvbody.AuthorId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		response.SendErrorResponse(w, response.InternalFaults)
	} else {
		response.SendNormalResponse(w, string(resp), 201)
	}

}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}

	uname := p.ByName("username")
	vs, err := model.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListAllvideos: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}

	vsi := &model.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(vsi); err != nil {
		response.SendErrorResponse(w, response.InternalFaults)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}

}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	cbody := &model.NewComment{}
	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Printf("%s", err)
		response.SendErrorResponse(w, response.RequestBodyParseFailedError)
		return
	}

	vid := p.ByName("vid-id")
	if err := model.AddNewComment(vid, cbody.AuthorId, cbody.Content); err != nil {
		log.Printf("Error in PostComment: %s", err)
		response.SendErrorResponse(w, response.DBError)
	} else {
		response.SendNormalResponse(w, "ok", 201)
	}

}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !ValidateUser(w, r) {
		return
	}

	vid := p.ByName("vid-id")
	cm, err := model.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		response.SendErrorResponse(w, response.DBError)
		return
	}

	cms := &model.Comments{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		response.SendErrorResponse(w, response.InternalFaults)
	} else {
		response.SendNormalResponse(w, string(resp), 200)
	}
}
