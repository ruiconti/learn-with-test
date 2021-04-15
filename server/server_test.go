package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Requirements:
// Create a webserver where users can track how many games players have won.
// Expected routes:
//	- GET	/players/{name}
//		Returns a number indicating the total numbers of wins.
//	- POST  /players/{name}
//		Record a win, incrementing for every subsequent POST
//
// Strategy:
//	1. Implement a `PlayerStore` interface to store players' scores. As an
//		interface, it makes easy to later adapt a persitent storage to it.
//		Also, through a simple `PlayerStoreStub`, we can get feedback on
//		expected behaviour.
//	2. We could spy on `PlayerStoreStub` to check whether POSTS are properly
//		fulfilled

func TestStoreGetScore(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"caddy": 20,
			"mary":  10,
		},
		winCalls: []string{},
	}
	server := &PlayerServer{store}

	t.Run("returns Caddy's score", func(t *testing.T) {
		req := newGetScoreRequest("caddy")
		response := httptest.NewRecorder()
		// NewRecorder is a Response stub.

		server.ServeHTTP(response, req)

		assertResponseBodyString(t, "20", response)
		assertResponseStatus(t, http.StatusOK, response)

	})

	t.Run("returns Mary's score", func(t *testing.T) {
		req := newGetScoreRequest("mary")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, req)

		assertResponseBodyString(t, "10", response)
		assertResponseStatus(t, http.StatusOK, response)
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		req := newGetScoreRequest("jonas")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, req)

		assertResponseStatus(t, http.StatusNotFound, response)
	})
}

func TestStoreWriteScore(t *testing.T) {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}
	server := &PlayerServer{store}

	t.Run("it increases store on POST", func(t *testing.T) {
		player := "barbra"
		reqPost := newPostWinsRequest(player)
		responsePost := httptest.NewRecorder()

		server.ServeHTTP(responsePost, reqPost)

		assertResponseStatus(t, http.StatusOK, responsePost)

		got := len(store.winCalls)
		want := 1

		if got != want {
			t.Errorf("got %d calls to RecordWin, wanted %d", got, want)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, got %v want %v", store.winCalls[0], player)
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	r, _ := http.NewRequest(GET, fmt.Sprintf("/players/%s", name), nil)
	return r
}

func newPostWinsRequest(name string) *http.Request {
	r, _ := http.NewRequest(POST, fmt.Sprintf("/players/%s", name), nil)
	return r
}

func assertResponseBodyString(t testing.TB, want string, response *httptest.ResponseRecorder) {
	t.Helper()
	got := response.Body.String()

	if got != want {
		t.Errorf("got response body string %v, wanted %v", got, want)
	}
}

func assertResponseStatus(t testing.TB, want int, response *httptest.ResponseRecorder) {
	t.Helper()
	got := response.Code

	if got != want {
		t.Errorf("got response status %v, wanted %v", got, want)
	}
}
