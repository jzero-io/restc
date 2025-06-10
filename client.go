package restc

import (
	"net/http"
	"net/url"
	"time"
)

type Client interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request

	GetHeader() http.Header
	SetHeader(header http.Header)
}

type Opt func(client *client) error

type client struct {
	protocol string
	addr     string
	port     string

	retryTimes int
	retryDelay time.Duration

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func (r *client) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *client) Post() *Request {
	return r.Verb("POST")
}

func (r *client) Get() *Request {
	return r.Verb("GET")
}

func (r *client) GetHeader() http.Header {
	return r.headers
}

func (r *client) SetHeader(header http.Header) {
	r.headers = header
}

func NewClient(target string, ops ...Opt) (Client, error) {
	c := &client{}

	// parse url
	parse, err := url.Parse(target)
	if err != nil {
		return nil, err
	}
	c.protocol = parse.Scheme
	c.addr = parse.Hostname()
	c.port = parse.Port()

	for _, op := range ops {
		if err = op(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
