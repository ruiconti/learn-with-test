package main

import (
	"context"
	"fmt"
	"net/http"
)

// Scenario:
// - Webserver that fetches for some data and returns the response
// - That fetching might take a while
// - What happens when a user cancels the request before the data can be
//	 retrieved?

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}
