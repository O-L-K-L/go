package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var port = flag.String("port", ":5050", "server port")

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/tasks", taskHandler)

	log.Fatal(http.ListenAndServe(*port, nil))
}

type Feature struct {
	Name string `json:"name"`
}

type Task struct {
	Feature Feature `json:"feature"`
	Name    string  `json:"name"`
	OrigEst string  `json:"originalEstimatedTime"`
	CurEst  string  `json:"currentEstimatedTime"`
	Elapsed string  `json:"elapsedTime"`
}

var features = []Feature{
	{Name: "HTTP Server"},
}

var tasks = []Task{
	{
		Feature: features[0],
		Name:    "Task CRUD",
		OrigEst: "2h",
		CurEst:  "2h",
		Elapsed: "30m",
	},
}

func taskHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		s, err := json.Marshal(tasks)

		if err != nil {
			w.WriteHeader(400)
			return
		}

		w.Header().Add("content-type", "application/json")
		w.Write(s)
	} else if req.Method == http.MethodPost {
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)

		if err != nil {
			w.WriteHeader(400)
			log.Println("failed to read body", err)
			return
		}

		var newTask Task

		if err = json.Unmarshal(body, &newTask); err != nil {
			w.WriteHeader(400)
			log.Println("failed to convert to json", err)
			return
		}

		tasks = append(tasks, newTask)
		w.WriteHeader(204)
	}
}
