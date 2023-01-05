package cache

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type ReadCloserCopy struct {
	r io.Reader
}

func (r *ReadCloserCopy) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

func (r *ReadCloserCopy) Close() error {
	return nil
}

type Transport struct {
	Next  http.RoundTripper
	cache map[string]*http.Response
}

func copyReaderCloser(r io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	all, err := io.ReadAll(r)
	if err != nil {
		log.Fatal("copy failed")
		return nil, nil
	}

	return &ReadCloserCopy{r: bytes.NewReader(all)}, &ReadCloserCopy{r: bytes.NewReader(all)}
}

func (c *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	if c.cache == nil {
		log.Println("initialize cache transport")
		c.cache = make(map[string]*http.Response)
	}

	u := request.URL.String()
	cachedResponse, isCached := c.cache[u]
	if isCached {
		log.Printf("CACHED: %v\n", u)
		return cachedResponse, nil
	}

	log.Printf("NOT CACHED: %v\n", u)
	response, err := c.Next.RoundTrip(request)
	if err != nil {
		return nil, err
	}

	c.cache[u] = response
	return response, nil
}
