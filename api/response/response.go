package response

import (
	"net/http"
	"encoding/json"
	"io"
)

func SendErrorResponse(w http.ResponseWriter, response ErrorResponse) {
	w.WriteHeader(response.HttpSC)
	bytes, _ := json.Marshal(&response.Error)
	io.WriteString(w, string(bytes))
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
