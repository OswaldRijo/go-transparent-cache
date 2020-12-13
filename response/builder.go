package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const StatusOk = "OK"
const RequestSuccessfully = "Request was resolved successfully"
const StatusError = "Error"

type StatusCode int

type Status string
type Content interface {}
type Data interface {}

type Message struct {
	StatusCode
	Status
	Content
}

type okMessage struct {
	Status `json:"status"`
	Data   `json:"data"`
}

type errMessage struct {
	Status `json:"status"`
	Data   `json:"message"`
}

func setInternalErrorServer(w http.ResponseWriter)  {
	fmt.Printf("error processing response json")
	var errMessage = errMessage{"Error", "error"}
	var errBody, _ =json.Marshal(errMessage)
	(w).WriteHeader(http.StatusInternalServerError)
	(w).Write(errBody)

}
func SetOkResponse(data *Message, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseBytes, errJson := json.Marshal(okMessage{data.Status, data.Content})

	if errJson != nil {
		setInternalErrorServer(w)
	}
	w.WriteHeader(int(data.StatusCode))
	w.Write(responseBytes)
}

func SetErrorResponse(data *Message, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseBytes, errJson := json.Marshal(errMessage{data.Status, data.Content})

	if errJson != nil {
		setInternalErrorServer(w)
	}
	w.WriteHeader(int(data.StatusCode))
	w.Write(responseBytes)
}