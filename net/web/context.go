package web

import (
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Keys     map[string]interface{}
	done     bool
}

// Set is used to store a new key/value pair exclusivelly for this context.
// It also lazy initializes c.Keys if it was not used previously.
func (c *Context) Set(key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *Context) Get(key string) (value interface{}, exists bool) {
	if c.Keys != nil {
		value, exists = c.Keys[key]
	}
	return
}

func (c *Context) Done() {
	c.done = true
}
