package handlers

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	messagesMutex sync.RWMutex
	messages      = make(map[string][]string)
)

const (
	timeoutDuration = 30 * time.Second
)

func MessageReceiver(w http.ResponseWriter, r *http.Request, topic string) {
	// Set response headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	// Set response status
	w.WriteHeader(http.StatusOK)

	// Send messages in SSE format
	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	messagesMutex.Lock()
	messagesByTopic := messages[topic]
	messagesMutex.Unlock()

	// Set up a timeout mechanism to disconnect clients
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	notify := w.(http.CloseNotifier).CloseNotify()
	done := make(chan bool)

	go func() {
		<-notify
		done <- true
	}()

	for _, msg := range messagesByTopic {
		fmt.Fprintf(w, "%s \n", msg)
		flusher.Flush()

		// Reset the timer
		if !timeout.Stop() {
			<-timeout.C
		}

		timeout.Reset(timeoutDuration)

		select {
		case <-notify:
			done <- true
			return
		case <-timeout.C:
			msg := "event: timeout \ndata: " + timeoutDuration.String() + "\n\n"
			fmt.Fprintf(w, msg)
			flusher.Flush()
			done <- true
			return
		default:
		}
	}

	select {
	case <-notify:
		<-done
		return
	case <-timeout.C:
		msg := "event: timeout \ndata: " + timeoutDuration.String() + "\n\n"
		fmt.Fprintf(w, msg)
		flusher.Flush()
		return
	case <-done:
		return
	}
}
