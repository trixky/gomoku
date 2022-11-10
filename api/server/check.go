package server

import (
	"fmt"

	"encoding/json"
	"net/http"

	dt "github.com/trixky/gomoku/doubleThree"
	"github.com/trixky/gomoku/models"
)

type RequestCheckData struct {
	Goban  string `json:"goban"`
	X      int    `json:"X"`
	Y      int    `json:"Y"`
	Player int    `json:"Player"`
}

type ResponseCheckData struct {
	DoubleThree bool   `json:"DoubleThree"`
	NbCaptured  int    `json:"NbCaptured"`
	Goban       string `json:"goban"`
	Win         bool   `json:"win"`
}

func check(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Decode the data from JSON
	decoder := json.NewDecoder(r.Body)
	data := models.RequestNextData{}

	if err := decoder.Decode(&data); err != nil {
		// If JSON unmarshalling failed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Compute a new context from the data
	context, err := data.ComputeContext()

	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doubleThree, nb, goban := dt.CheckDoubleThree(context.Goban, context.State.LastMove.Position, context.State.LastMove.Player)

	response := ResponseCheckData{}

	dt.IsWin(&context)
	response.DoubleThree = doubleThree
	response.NbCaptured = nb
	response.Goban = goban.ToString()
	response.Win = false
	if (context.State.LastMove.Player == false &&
		context.State.PlayersInfo.Player_1.Win == true) ||
		context.State.PlayersInfo.Player_2.Win == true {
		response.Win = true
	}

	marshalled, err := json.Marshal(response)

	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	marshalledStr := string(marshalled)

	fmt.Fprintf(w, marshalledStr)
}
