package internal

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	Target  *url.URL
	Method  string
	Timeout time.Duration
	Payload []byte
}

func (r Request) Size() int { return len(r.Payload) }

func (r Request) toRequest(ctx context.Context) (*http.Request, error) {
	reqCtx := context.WithTimeout(ctx, r.Timeout)
	body := bytes.NewBuffer(r.Payload)
	return http.NewRequestWithContext(reqCtx, r.Method, r.Target.String(), body)
}

type Result struct {
	Timestamp                   time.Time
	Duration                    time.Duration
	StatusCode, InSize, OutSize int
	err                         error
}

func Run(ctx context.Context, req Request) (result Result) {
	hreq, err := req.toRequest(ctx)
	if err != nil {
		result.err = err
		return
	}
	result.Timestamp = time.Now()
	res, err := http.DefaultClient.Do(hreq)
	result.Duration = time.Now().Sub(result.Timestamp)
	if err != nil {
		result.err = err
		return
	}
	defer res.Body.Close()
	result.StatusCode = r.StatusCode
	result.InSize, result.OutSize = r.Size(), int(res.ContentLength)
	return
}

type User struct {
	NumOfRequests int
}

func (u User) Start(ctx context.Context, req Request) []Result {
	results := make([]Result, 0, u.NumOfRequests)
	for i := 0; i < u.NumOfRequests; i++ {
		results = append(results, Run(ctx, req))
	}
	return results
}
