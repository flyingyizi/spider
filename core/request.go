package core

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request is the representation of a HTTP request made by a Collector
type Request struct {
	// URL is the parsed URL of the HTTP request
	URL *url.URL
	// Headers contains the Request's HTTP headers
	Headers *http.Header
	// Ctx is a context between a Request and a Response
	Ctx *Context
	// Depth is the number of the parents of the request
	Depth int
	// Method is the HTTP method of the request
	Method string
	// Body is the request body which is used on POST/PUT requests
	Body io.Reader
	// ResponseCharacterencoding is the character encoding of the response body.
	// Leave it blank to allow automatic character encoding of the response body.
	// It is empty by default and it can be set in OnRequest callback.
	ResponseCharacterEncoding string
	// ID is the Unique identifier of the request
	ID        uint32
	collector ICollector
	abort     bool
	baseURL   *url.URL
	// ProxyURL is the proxy address that handles the request
	ProxyURL string
}

type serializableRequest struct {
	URL     string
	Method  string
	Body    []byte
	ID      uint32
	Ctx     map[string]interface{}
	Headers http.Header
}

// New creates a new request with the context of the original request
func (r *Request) New(method, URL string, body io.Reader) (*Request, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	return &Request{
		Method:    method,
		URL:       u,
		Body:      body,
		Ctx:       r.Ctx,
		Headers:   &http.Header{},
		ID:        r.collector.NewID(), //atomic.AddUint32(&r.collector.requestCount, 1)
		collector: r.collector,
	}, nil
}

// NewRequest creates a new request
func NewRequest(method string, u *url.URL, body io.Reader,
	ctx *Context, header *http.Header,
	RequestID uint32, depth int, collector ICollector) *Request {

	return &Request{
		Method:    method,
		URL:       u,
		Body:      body,
		Ctx:       ctx,
		Headers:   header,
		ID:        RequestID,
		collector: collector,
		Depth:     depth,
	}

}

// Abort cancels the HTTP request when called in an OnRequest callback
func (r *Request) Abort() {
	r.abort = true
}
func (r *Request) IsAbort() bool {
	return r.abort
}

//SetBaseURL value
func (r *Request) SetBaseURL(u *url.URL) {
	r.baseURL = u
}

// AbsoluteURL returns with the resolved absolute URL of an URL chunk.
// AbsoluteURL returns empty string if the URL chunk is a fragment or
// could not be parsed
func (r *Request) AbsoluteURL(u string) string {
	if strings.HasPrefix(u, "#") {
		return ""
	}
	var base *url.URL
	if r.baseURL != nil {
		base = r.baseURL
	} else {
		base = r.URL
	}
	absURL, err := base.Parse(u)
	if err != nil {
		return ""
	}
	absURL.Fragment = ""
	if absURL.Scheme == "//" {
		absURL.Scheme = r.URL.Scheme
	}
	return absURL.String()
}

// Visit continues Collector's collecting job by creating a
// request and preserves the Context of the previous request.
// Visit also calls the previously provided callbacks
func (r *Request) Visit(URL string) error {
	return r.collector.Scrape(r.AbsoluteURL(URL), "GET", r.Depth+1, nil, r.Ctx, nil, true)
}

// Post continues a collector job by creating a POST request and preserves the Context
// of the previous request.
// Post also calls the previously provided callbacks
func (r *Request) Post(URL string, requestData map[string]string) error {
	return r.collector.Scrape(r.AbsoluteURL(URL), "POST", r.Depth+1, createFormReader(requestData), r.Ctx, nil, true)
}

// PostRaw starts a collector job by creating a POST request with raw binary data.
// PostRaw preserves the Context of the previous request
// and calls the previously provided callbacks
func (r *Request) PostRaw(URL string, requestData []byte) error {
	return r.collector.Scrape(r.AbsoluteURL(URL), "POST", r.Depth+1, bytes.NewReader(requestData), r.Ctx, nil, true)
}

// PostMultipart starts a collector job by creating a Multipart POST request
// with raw binary data.  PostMultipart also calls the previously provided.
// callbacks
func (r *Request) PostMultipart(URL string, requestData map[string][]byte) error {
	boundary := randomBoundary()
	hdr := http.Header{}
	hdr.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	hdr.Set("User-Agent", r.collector.UserAgent())
	return r.collector.Scrape(r.AbsoluteURL(URL), "POST", r.Depth+1, createMultipartReader(boundary, requestData), r.Ctx, hdr, true)
}

// Retry submits HTTP request again with the same parameters
func (r *Request) Retry() error {
	return r.collector.Scrape(r.URL.String(), r.Method, r.Depth, r.Body, r.Ctx, *r.Headers, false)
}

// Do submits the request
func (r *Request) Do() error {
	return r.collector.Scrape(r.URL.String(), r.Method, r.Depth, r.Body, r.Ctx, *r.Headers, !r.collector.AllowURLRevisit())
}

// Marshal serializes the Request
func (r *Request) Marshal() ([]byte, error) {
	ctx := make(map[string]interface{})
	if r.Ctx != nil {
		r.Ctx.ForEach(func(k string, v interface{}) interface{} {
			ctx[k] = v
			return nil
		})
	}
	var err error
	var body []byte
	if r.Body != nil {
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
	}
	sr := &serializableRequest{
		URL:    r.URL.String(),
		Method: r.Method,
		Body:   body,
		ID:     r.ID,
		Ctx:    ctx,
	}
	if r.Headers != nil {
		sr.Headers = *r.Headers
	}
	return json.Marshal(sr)
}
