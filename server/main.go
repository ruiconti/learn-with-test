package main

import (
	"log"
	"net/http"
)

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{scores: map[string]int{}}
}

type InMemoryPlayerStore struct {
	scores map[string]int
}

func (i *InMemoryPlayerStore) GetScore(name string) (int, error) {
	score, found := i.scores[name]
	if !found {
		return 0, PlayerNotFoundErr
	}
	return score, nil
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.scores[name]++
}

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
