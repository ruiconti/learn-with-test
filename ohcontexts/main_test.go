package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("stub got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		s.cancelled = true
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func (s *StubStore) assertCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *StubStore) assertNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("server responds accordingly", func(t *testing.T) {
		data := "hi!"
		store := &StubStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s' wanted '%s'", response.Body.String(), data)
		}

		store.assertNotCancelled()
	})

	t.Run("user cancels during fetch", func(t *testing.T) {
		data := "long running call"
		store := &StubStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		ctx := request.Context()

		cancellingCtx, cancel := context.WithCancel(ctx)
		// Returns a copy of the parent with a new Done channel and a cancel fn

		time.AfterFunc(5*time.Millisecond, cancel)
		// Calls cancel function after 5ms
		request = request.WithContext(cancellingCtx)
		// Updates parent request with new derived context

		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		store.assertCancelled()
	})
}
