package main

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status `json:"status"`
}

func main() {
}
