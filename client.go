package restc

import (
	"net/http"
	"time"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request

	GetHeader() http.Header
	SetHeader(header http.Header)
}

type Opt func(client *RestClient) error

type RestClient struct {
	protocol string
	addr     string
	port     string

	retryTimes int
	retryDelay time.Duration

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func (r *RestClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *RestClient) Post() *Request {
	return r.Verb("POST")
}

func (r *RestClient) Get() *Request {
	return r.Verb("GET")
}

func (r *RestClient) GetHeader() http.Header {
	return r.headers
}

func (r *RestClient) SetHeader(header http.Header) {
	r.headers = header
}

func New(ops ...Opt) (*RestClient, error) {
	c := &RestClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
