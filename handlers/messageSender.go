package handlers

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	messageId int
)

func MessageSender(w http.ResponseWriter, r *http.Request, topic string) {
	// Read message from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add message to the messages map
	messageId++
	id := strconv.Itoa(messageId)

	messages[topic] = append(messages[topic], "id: "+id+" \nevent: msg \ndata: "+string(body)+"\n\n")

	// Send 204 No Content response
	w.WriteHeader(http.StatusNoContent)
	return
}
