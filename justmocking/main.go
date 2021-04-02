package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Requirements:
// 1. Write a function that prints:
// 3
// 2
// 1
// Go!
// Each line should be printed after 1 sec.
func main() {
	sleepSeconds := 1
	count := 5

	StartCountdown(count, sleepSeconds)
}

func StartCountdown(count, secondsToSleep int) {
	duration := time.Second * time.Duration(secondsToSleep)
	sleeper := &RealSleeper{time.Sleep, duration}
	Countdown(os.Stdout, count, sleeper)
}

type Sleeper interface {
	Sleep()
}

type RealSleeper struct {
	sleep    func(time.Duration)
	duration time.Duration
}

func (r *RealSleeper) Sleep() {
	r.sleep(r.duration)
}

type FakeSleeper struct {
	durationSlept time.Duration
	Calls         int
}

func (s *FakeSleeper) Sleep() {
	s.Calls++
}

const write = "write"
const sleep = "sleep"

type FakeCountdownOperations struct {
	Calls []string
}

func (c *FakeCountdownOperations) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *FakeCountdownOperations) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func Countdown(out io.Writer, start int, sleeper Sleeper) {
	for i := start; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}
