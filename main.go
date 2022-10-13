package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status `json:"status"`
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
