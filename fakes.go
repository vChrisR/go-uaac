package uaa

type FakeUaac struct {
	req          *Request
	responseBody []byte
	err          error
}

func (c *FakeUaac) NewRequestReturns(r *Request) {
	c.req = r
}

func (c *FakeUaac) NewExecuteRequest(body []byte, err error) {
	c.responseBody = body
	c.err = err
}

func (c *FakeUaac) NewRequest(method, path string) *Request {
	return c.req
}
func (c *FakeUaac) ExecuteRequest(r *Request) ([]byte, error) {
	return c.responseBody, c.err
}
func (c *FakeUaac) ExecuteAndUnmarshall(r *Request, target interface{}) {
	panic("fakeUaac.ExecuteAndUnmarshall not yet implemented")
}
