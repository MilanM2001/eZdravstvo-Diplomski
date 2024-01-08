package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func jsonResponse(object interface{}, w http.ResponseWriter) {
	resp, err := json.Marshal(object)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getIDFromReqAsPrimitive(writer http.ResponseWriter, req *http.Request) (primitive.ObjectID, error) {
	vars := mux.Vars(req)
	id := vars["id"]

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Convert to Primitive error")
		writer.WriteHeader(http.StatusBadRequest)
		return primitive.NilObjectID, err
	}

	return objectID, nil
}
