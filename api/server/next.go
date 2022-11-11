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

	context.Goban.ComputeGlobalProximity(context.Options.ProximityThreshold, context.Options.ProximityRadius, context.Options.ProximityShape, context.Options.SuspicionRadius, context.State.LastMove)

	childs := []models.Context{}
	var child_channel chan *models.Context
	heuristic_goban := models.HeuristicGoban{}
	var best_child models.Context

	// context.Goban.PrintProximity(&context.Options.ProximityThreshold)

	blocked := context.Goban.IsBlocked(context.Options.ProximityThreshold)

	if blocked {
		goto choose
	}

	// context.Options.Print()
	// context.State.Print()

	child_channel = make(chan *models.Context, 1)

	childs = logic.Negamax(&context, child_channel)

	<-child_channel

	heuristic_goban.Print()

	heuristic_goban.Compute(childs)

choose:
	if len(childs) == 0 {
		best_child = logic.Random(&context, !blocked)
	} else {
		best_child = childs[len(childs)-1]
		for _, child := range childs {
			if child.State.Beta > best_child.State.Beta {
				best_child = child
			}
		}
	}

	// best_child.Print()

	elapsed_time := time.Now().Sub(context.Start).Milliseconds()

	_json, err := best_child.ToJSON(elapsed_time, heuristic_goban)

	if err != nil {
		// If the context computation failed
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, _json)
}
