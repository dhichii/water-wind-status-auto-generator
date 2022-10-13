package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status `json:"status"`
}

type Level struct {
	Water string
	Wind  string
}

type Response struct {
	Status
	Level
}

func updateData() {
	for {
		var data = Data{Status: Status{}}
		statusMin := 1
		statusMax := 30

		data.Status.Water = rand.Intn(statusMax-statusMin) + statusMin

		data.Status.Wind = rand.Intn(statusMax-statusMin) + statusMin

		b, err := json.MarshalIndent(&data, "", "  ")

		if err != nil {
			log.Fatalln("error while marshalling json data  =>", err.Error())
		}

		err = ioutil.WriteFile("data.json", b, 0644)

		if err != nil {
			log.Fatalln("error while writing value to data.json file  =>", err.Error())
		}
		time.Sleep(time.Second * 15)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go updateData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("index.html")

		b, err := ioutil.ReadFile("data.json")
		if err != nil {
			fmt.Fprint(w, "read file error")
			return
		}

		var data = Data{Status: Status{}}
		if err := json.Unmarshal(b, &data); err != nil {
			fmt.Fprint(w, "marshalling error")
		}

		var response = Response{Status: data.Status}
		response.Level.Water = evaluateWater(data.Status.Water)
		response.Level.Wind = evaluateWind(data.Status.Wind)

		tpl.ExecuteTemplate(w, "index.html", response)

	})

	http.ListenAndServe(":8080", nil)
}

func evaluateWater(status int) string {
	if status > 8 {
		return "bahaya"
	}

	if status > 5 {
		return "siaga"
	}

	return "aman"
}

func evaluateWind(status int) string {
	if status > 15 {
		return "bahaya"
	}

	if status > 7 {
		return "siaga"
	}

	return "aman"
}
