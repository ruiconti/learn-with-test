package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	GET  = http.MethodGet
	POST = http.MethodPost
)

var PlayerNotFoundErr = errors.New("Player not found")

// Stubs for testing
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

// Stub's fake method for retrieval
func (s *StubPlayerStore) GetScore(name string) (int, error) {
	score, found := s.scores[name]
	if !found {
		return 0, PlayerNotFoundErr
	}
	return score, nil
}

// Stub's fake method for writing to enable spying
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// Actual PlayerStore Interface
type PlayerStore interface {
	GetScore(name string) (int, error)
	RecordWin(name string)
}

// Fake for actual retrieval
func GetScore(playerName string) (int, error) {
	var score int

	switch playerName {
	case "caddy":
		score = 20
	case "mary":
		score = 10
	default:
		return 0, PlayerNotFoundErr
	}
	return score, nil
}

type PlayerServer struct {
	store PlayerStore
}

func (ps *PlayerServer) handleGetScore(w http.ResponseWriter, name string) {
	score, err := ps.store.GetScore(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (ps *PlayerServer) handleRecordWin(w http.ResponseWriter, name string) {
	ps.store.RecordWin(name)
	w.WriteHeader(http.StatusOK)
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case POST:
		ps.handleRecordWin(w, playerName)
	case GET:
		ps.handleGetScore(w, playerName)
	}
}

// func PlayerServer(w http.ResponseWriter, request *http.Request) {
// 	playerName := strings.TrimPrefix(request.URL.Path, "/players/")
// 	// Equals:
// 	// urlPath := request.URL.Path
// 	// playerName := urlPath[9:]
//
// 	score := GetPlayerScore(playerName)
//
// 	fmt.Fprint(w, score)
// 	// Equals:
// 	// val := []byte("20")
// 	// w.Write(val)
// }
