package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createDelayingServer(delay time.Duration) *httptest.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}

	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestRacer(t *testing.T) {
	t.Run("returns which URL responds faster", func(t *testing.T) {
		fastServer := createDelayingServer(0 * time.Millisecond)
		slowServer := createDelayingServer(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		got, err := Racer(slowURL, fastURL)
		want := fastURL

		if err != nil {
			t.Errorf("didn't want but got an error %v", err.Error())
		}

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}

	})
}

func TestConfigurableRacer(t *testing.T) {
	t.Run("timeouts if none of URLs responds in 10s", func(t *testing.T) {
		server := createDelayingServer(25 * time.Millisecond)

		defer server.Close()

		timeout := time.Millisecond * 20
		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil {
			t.Errorf("expected a timeout error")
		}

		want := TimeoutError(server.URL, server.URL)

		if err.Error() != want.Error() {
			t.Errorf("got error %v but wanted %v", err.Error(), want.Error())
		}
	})
}
