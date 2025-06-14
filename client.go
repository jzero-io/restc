package restc

import (
	"net/http"
	"net/url"
	"sync"
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
	lock     *sync.RWMutex
	protocol string
	addr     string
	port     string

	retryTimes int
	retryDelay time.Duration

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client

	// middleware
	beforeRequest []RequestMiddleware
}

type RequestMiddleware func(Client, *Request) error

func (c *client) requestMiddlewares() []RequestMiddleware {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.beforeRequest
}

func (c *client) executeRequestMiddlewares(req *Request) (err error) {
	for _, f := range c.requestMiddlewares() {
		if err = f(c, req); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

func (c *client) Post() *Request {
	return c.Verb("POST")
}

func (c *client) Get() *Request {
	return c.Verb("GET")
}

func (c *client) GetHeader() http.Header {
	return c.headers
}

func (c *client) SetHeader(header http.Header) {
	c.headers = header
}

func NewClient(target string, ops ...Opt) (Client, error) {
	c := &client{
		lock: &sync.RWMutex{},
	}

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
