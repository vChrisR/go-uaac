package uaa

import "github.com/dave-malone/oauth"

type FakeUaac struct {
	req          *oauth.Request
	responseBody []byte
	err          error
}

func (c *FakeUaac) NewRequestReturns(r *oauth.Request) {
	c.req = r
}

func (c *FakeUaac) NewExecuteRequest(body []byte, err error) {
	c.responseBody = body
	c.err = err
}

func (c *FakeUaac) NewRequest(method, path string) *oauth.Request {
	return c.req
}
func (c *FakeUaac) ExecuteRequest(r *oauth.Request) ([]byte, error) {
	return c.responseBody, c.err
}
func (c *FakeUaac) ExecuteAndUnmarshall(r *oauth.Request, target interface{}) {
	panic("fakeUaac.ExecuteAndUnmarshall not yet implemented")
}
