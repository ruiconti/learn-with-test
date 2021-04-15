package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrieving(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Jonas"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinsRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertResponseStatus(t, http.StatusOK, response)
	assertResponseBodyString(t, "3", response)
}
