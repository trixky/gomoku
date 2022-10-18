package server

import (
	"fmt"
	"log"
	"time"
	"strings"
	"net/http"
	"io/ioutil"

)

func resolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	if r.Method != "POST" {
		// If the request method is not POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}
	
	fmt.Fprintf(w, strToFront)
}
