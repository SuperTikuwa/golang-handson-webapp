package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SuperTikuwa/golang-handson-todo/dbctl"
	"github.com/SuperTikuwa/golang-handson-todo/model"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT")
	switch r.Method {
	case http.MethodPost:
		postHandler(w, r)
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPut:
		putHandler(w, r)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var requestJSON model.TaskPostRequest

	if err := json.Unmarshal(body, &requestJSON); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dbTask model.Task
	dbTask.Title = requestJSON.Title
	dbTask.Description = requestJSON.Description
	dbTask.Date = requestJSON.Date

	if err := dbctl.InsertNewTask(dbTask); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(dbTask)
	w.WriteHeader(http.StatusOK)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	response, err := dbctl.SelectAllTasks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonBytes))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var putRequest model.TaskPutRequest
	if err := json.Unmarshal(body, &putRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := dbctl.UpdateToDone(putRequest.ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
