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
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	userCredential := &model.UserCredential{}
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
	userCredential := &model.UserCredential{}
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
