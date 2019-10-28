package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper interface
type Sleeper interface {
	Sleep()
}

// SpySleeper struct
type SpySleeper struct {
	Calls int
}

// Sleep method for SpySleeper
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// SpyTime struct
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep method for SpyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// ConfigurableSleeper struct
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep method for ConfigurableSleeper
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// CountdownOperationsSpy struct
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep method for CountdownOperationsSpy
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write method for CountdownOperationsSpy
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const countdownStart = 3
const finalWord = "Go!"
const sleep = "sleep"
const write = "write"

// Countdown function
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
