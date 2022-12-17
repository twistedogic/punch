package internal

import (
	"net/url"
	"time"
)

type ExecMode uint

const (
	PARALLEL ExecMode = iota
	SERIAL
)

type Config struct {
	Mode             ExecMode
	Metrics, Profile Scrape
	Targets          []Target
}

type Scrape struct {
	scrapeTarget   *url.URL
	scrapeInterval time.Duration
}

// Target defines the test scenerio
// It can end by specifying an end duration or
// error rate threshold is reached
type Target struct {
	targetURL            *url.URL
	method               string
	payload              []byte
	expectedStatusCode   int
	timeout, runDuration time.Duration
	threshold            float64
	bucket               []int
}
