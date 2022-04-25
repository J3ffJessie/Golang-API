package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"errors"
)

var (
	//go:embed testData.json
	testData []byte
	running = true
)

func main() {

	fmt.Println("starting server....")
	http.HandleFunc(":8080", routeHandler)
}

type reqBody struct {
	Message map[string]interface{} `json:"message"`
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handlePOST(w, r)
	case "GET":
		handleGET(w, r)

	}

}
//TODO: finish logic for post functionality
	//TODO: POST calls do not change data already in testData.json file only adds to it
	//LIMIT: we are not persisting the data
func handlePOST(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var db reqBody
	_=json.Unmarshal(testData, & db)
	var req map[string]interface{}

    buf := new(bytes.Buffer)

    l, err := buf.ReadFrom(r.Body)
	fmt.Printf("%v", err)
	fmt.Printf("%v", l)
	

    err = json.Unmarshal(buf.Bytes(), &req)
	if err != nil{
		fmt.Println(err.Error())
	}

	for k, v := range req {
		if db.Message[k] == nil {
			db.Message[k] = v
		}

	}

	fmt.Printf("Check out the data after you posted!!!")
	json.NewEncoder(w).Encode(db.Message) //TODO:encode response here
}

func handleGET(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req reqBody
	_=json.Unmarshal(testData, & req)

	fmt.Println("you Got it!!!!")
	json.NewEncoder(w).Encode(req) //TODO:encode response here
}

// declare a function with multiple returns
func multipleReturns() (string, error) {
	return "string", errors.New("")
}