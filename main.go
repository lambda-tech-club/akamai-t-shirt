package main

import (
	"fmt"
	"html"
	// Some packages starting with "lo" should be used in the hidden method, I guess it will be doStuff.
	// "lo" -> "log" ?
	"net/http"
	"strconv"
	"time"
)

type ControlMessage struct {
	Target string
	Count  uint64
}

func main() {
	controlChannel := make(chan ControlMessage)
	workerCompleteChan := make(chan bool)
	statusPollChannel := make(chan chan bool)

	workerActive := false
	go admin(controlChannel, statusPollChannel)
	for {
		select {
		case respChan := <-statusPollChannel:
			respChan <- workerActive
		case msg := <-controlChannel:
			workerActive = true
			go doStuff(msg, workerCompleteChan)
		case status := <-workerCompleteChan:
			workerActive = status
		}
	}
}

func admin(cc chan ControlMessage, statusPollChannel chan chan bool) {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// For what?
		// hostTokens := strings.Split(r.Host, "?")
		err := r.ParseForm()
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		count, err := strconv.ParseUint(r.FormValue("count"), 10, 32)
		if err != nil {
			// Bad request?
			fmt.Fprint(w, err.Error())
			return
		}

		target := r.FormValue("target")
		msg := ControlMessage{
			Target: target,
			Count:  count,
		}
		cc <- msg

		fmt.Fprintf(w, "Control message issued for Target: %s, Count: %d",
			html.EscapeString(r.FormValue("target")),
			count,
		)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		reqChan := make(chan bool)
		statusPollChannel <- reqChan
		timeout := time.After(time.Second)
		select {
		case result := <-reqChan:
			if result {
				fmt.Fprintf(w, "ACTIVE")
			} else {
				fmt.Fprintf(w, "INACTIVE")
			}
			// Why? Doesn't look like we need to return here
			return
		case <-timeout:
			fmt.Fprintf(w, "TIMEOUT")
		}
	})

	// It Seems to be an HTTP server, so start to listen.
	http.ListenAndServe(":3000", nil)
}

// Should be some process for UNIQLO.
// Or, for PEACE?
func doStuff(msg ControlMessage, cc chan bool) {
	fmt.Printf("Target: %s, Count: %d \n", msg.Target, msg.Count)
	// Not very sure what we need to pass here.
	// But since no one would pass parameters to the method that they don't need,
	// and the worker becomes active after getting the message.
	// Just making worker inactive when completed.
	cc <- false
}
