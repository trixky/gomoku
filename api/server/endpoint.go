package server

import (
	"fmt"

	// "time"
	"encoding/json"
	// "strings"
	"net/http"

	"github.com/trixky/gomoku/models"
	// "io/ioutil"
)

func next(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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

	context.Print()

	context.Negamax()

	fmt.Fprintf(w, "fuck off")
}
