// Copyright (c) 2015 All Right Reserved, Improbable Worlds Ltd.

// Improbable custom handlers for stuff.

package echo

import (
	"bytes"
)

// WithMiddleware wraps the given handler directly with a set of middleware (in last to first order).
func HandlerWithMiddleware(h Handler, middlewares ...Middleware) HandlerFunc {
	wh := wrapHandler(h)
	// Chain middleware with handler in the end
	for i := len(middlewares) - 1; i >= 0; i-- {
		m := middlewares[i]
		wrappedM := wrapMiddleware(m)
		wh = wrappedM(wh)
	}
	return wh
}

// ParamMap returns a map of all parameters that result from the Path parameter expansion.
func (c *Context) ParamMap() map[string]string {
	out := make(map[string]string)
	for i, _ := range c.pnames {
		out[c.pnames[i]] = c.pvalues[i]
	}
	return out
}

// RenderWithContentType renders a template with data and sends the specified content type with
// response and status code. Templates can be registered using `Echo.SetRenderer()`.
func (c *Context) RenderWithContentType(code int, contentType string, name string, data interface{}) (err error) {
	if c.echo.renderer == nil {
		return RendererNotRegistered
	}
	buf := new(bytes.Buffer)
	if err = c.echo.renderer.Render(buf, name, data); err != nil {
		return
	}
	c.response.Header().Set(ContentType, contentType)
	c.response.WriteHeader(code)
	c.response.Write(buf.Bytes())
	return
}

// SetNotFoundHandler registers a custom handlers used when no route was found.
func (e *Echo) SetNotFoundHandler(h HandlerFunc) {
	e.notFoundHandler = h
}
