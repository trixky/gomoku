package server

import (
	"fmt"

	"encoding/json"
	"net/http"
)

type RequestCheckData struct {
	Goban  string `json:"goban"`
	X      int    `json:"X"`
	Y      int    `json:"Y"`
	Player int    `json:"Player"`
}

type ResponseCheckData struct {
	doubleThree bool   `json:"doubleThree"`
	nbCaptured  int    `json:"nbCaptured"`
	Goban       string `json:"goban"`
}

func check(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the data from JSON
	decoder := json.NewDecoder(r.Body)
	data := RequestCheckData{}

	if err := decoder.Decode(&data); err != nil {
		// If JSON unmarshalling failed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Compute a new context from the data
	// doubleThree, nb, goban := dt.CheckDoubleThree(m.Goban{}.Extract(data.Goban), m.Position{X: data.X, Y: data.Y}, data.Player)

	// response := ResponseCheckData{}

	// response.doubleThree = doubleThree
	// response.nbCaptured = nb
	// response.Goban = goban.ToString()

	// marshalled, err := json.Marshal(response)

	// if err != nil {
	// 	// If the context computation failed
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// marshalledStr := string(marshalled)

	// fmt.Fprintf(w, marshalledStr)
	fmt.Fprintf(w, "in progress")
}
