package view

import (
	"Golang-challenge/controller"
	"Golang-challenge/response"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	defer r.Body.Close()
	var controllerInstance = controller.NewController()
	itemCode, ok := vars["item_code"]
	if !ok {
		fmt.Printf("No item code on request \n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var price, err = controllerInstance.GetPrice(itemCode)

	if err != nil {
		fmt.Printf("error getting user : %v \n", err)
		errMessage:= response.Message{StatusCode: http.StatusBadRequest, Status: response.StatusError, Content: price}
		response.SetErrorResponse(&errMessage, w)
		return
	}

	message:= response.Message{StatusCode: http.StatusOK, Status: response.RequestSuccessfully, Content: serializePrice(itemCode, price)}
	response.SetOkResponse(&message, w)

}

func GetMany(w http.ResponseWriter, r *http.Request) {
	var bodyPrices = deserializeBody(r.Body)
	var controllerInstance = controller.NewController()
	var prices, err = controllerInstance.GetPrices(bodyPrices.ItemCodes...)

	if err != nil {
		fmt.Printf("error getting user : %v \n", err)
		errMessage:= response.Message{StatusCode: http.StatusBadRequest, Status: response.StatusError, Content: ""}
		response.SetErrorResponse(&errMessage, w)
		return
	}
	message:= response.Message{StatusCode: http.StatusOK, Status: response.RequestSuccessfully,
		Content: serializePrices(bodyPrices.ItemCodes, prices)}

	response.SetOkResponse(&message, w)
}
