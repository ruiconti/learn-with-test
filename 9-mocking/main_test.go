package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	count := 3
	duration := time.Second * 1

	t.Run("Expected printing behavior", func(t *testing.T) {
		buffer := bytes.Buffer{}
		sleeper := &FakeSleeper{duration, 0}
		Countdown(&buffer, count, sleeper)

		want := "3\n2\n1\nGo!"
		got := buffer.String()

		assertValues(t, got, want)
		assertValues(t, strconv.Itoa(sleeper.Calls), strconv.Itoa(count))

	})

	t.Run("Expected operations order", func(t *testing.T) {
		watcher := &FakeCountdownOperations{}
		Countdown(watcher, count, watcher)

		want := []string{
			write, sleep,
			write, sleep,
			write, sleep,
			write,
		}
		got := watcher.Calls

		if !reflect.DeepEqual(want, watcher.Calls) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRealSleeper(t *testing.T) {
	duration := time.Second * 5
	sleeper := &FakeSleeper{duration, 0}

	sleeper.Sleep()

	got := sleeper.durationSlept
	want := duration

	assertValues(t, got.String(), want.String())
}

func assertValues(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
