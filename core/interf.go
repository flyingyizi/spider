package core

import (
	"io"
	"net/http"
)

//ICollector ...
type ICollector interface {
	NewID() uint32 //atomic.AddUint32(&r.collector.requestCount, 1)
	UserAgent() string
	AllowURLRevisit() bool
	Scrape(u, method string, depth int, requestData io.Reader, ctx *Context, hdr http.Header, checkRevisit bool) error
}
