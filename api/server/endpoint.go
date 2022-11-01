package server

import (
	"fmt"
	"time"

	"encoding/json"
	"net/http"

	"github.com/trixky/gomoku/logic"
	"github.com/trixky/gomoku/models"
)

func next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	start_time := time.Now()

	// Decode the data from JSON
	decoder := json.NewDecoder(r.Body)
	data := models.RequestData{}

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

	context.Goban.ComputeGlobalProximity(context.Options.ProximityThreshold, context.Options.ProximityRadius, context.Options.ProximityShape)

	context.Print()

	// channel := make(chan *models.Context, 1)

	childs := logic.Negamax(&context, nil)

	heuristic_goban := models.HeuristicGoban{}

	heuristic_goban.Compute(childs)

	best_child := childs[len(childs)-1]
	for _, child := range childs {
		if child.State.HeuristicScore > best_child.State.HeuristicScore {
			best_child = child
		}
	}
	// best_child := <-channel

	elapsed_time := time.Now().Sub(start_time).Milliseconds()

	_json, err := best_child.ToJSON(elapsed_time, heuristic_goban)

	best_child.Print()

	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, _json)
}
