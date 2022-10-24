package server

import (
	"fmt"
	"log"
	// "time"
	"encoding/json"
	// "strings"
	"net/http"
	// "io/ioutil"

)

type Goban struct {
	board [19][19]byte
}

type Options struct {
	depth int
	timeout int
	selection_treshold int
}

type requestData struct {
	options Options
	goban Goban
}

type Cryptoresponse []struct {
	Name              string    `json:"options"`
}

func resolve(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("POST")
	decoder := json.NewDecoder(r.Body)
    var t requestData
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    log.Println(t.goban)
    log.Println(t.options)
	
	fmt.Fprintf(w, "fuck off")
}
