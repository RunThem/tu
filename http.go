package tu

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Http struct {
	mux *http.ServeMux
}

type Context struct {
	// origin objects
	Response http.ResponseWriter
	Request  *http.Request

	// Request info
	Path   string
	Method string

	// Response info
	StatusCode int
}

type HandleFunc func(ctx *Context)

type H map[string]any

func newContext(method string, path string, w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response:   w,
		Request:    r,
		Path:       path,
		Method:     method,
		StatusCode: 0,
	}
}

func NewHttp() *Http {
	return &Http{http.NewServeMux()}
}

func (mod *Http) Run(addr string) error {
	return http.ListenAndServe(addr, mod.mux)
}

func (mod *Http) addRoute(method string, path string, handleFunc HandleFunc) {
	mod.mux.HandleFunc(method+" "+path, func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context{
			Response:   w,
			Request:    r,
			Path:       path,
			Method:     method,
			StatusCode: 0,
		}

		handleFunc(ctx)
	})
}

func (mod *Http) Get(path string, handleFunc HandleFunc) {
	mod.addRoute("GET", path, handleFunc)
}

func (mod *Http) Post(path string, handleFunc HandleFunc) {
	mod.addRoute("POST", path, handleFunc)
}

func (mod *Http) Put(path string, handleFunc HandleFunc) {
	mod.addRoute("PUT", path, handleFunc)
}

func (mod *Http) Delete(path string, handleFunc HandleFunc) {
	mod.addRoute("DELETE", path, handleFunc)
}

func (mod *Http) Patch(path string, handleFunc HandleFunc) {
	mod.addRoute("PATCH", path, handleFunc)
}

func (mod *Http) Options(path string, handleFunc HandleFunc) {
	mod.addRoute("OPTIONS", path, handleFunc)
}

func (mod *Http) Head(path string, handleFunc HandleFunc) {
	mod.addRoute("HEAD", path, handleFunc)
}

// Status takes an integer and sets the response statue to the integer given.
func (mod *Context) Status(code int) {
	mod.StatusCode = code
	mod.Response.WriteHeader(code)
}

// Header gets the first value associated with the given key. If there are no values associated with
// the key. Get return "".
func (mod *Context) Header(key string) string {
	return mod.Request.Header.Get(key)
}

// Query method retrieves the form data.
func (mod *Context) Query(key string) string {
	return mod.Request.URL.Query().Get(key)
}

// SetHeader sets the header entries associated with key to the single element value. It replaces any existing values associated with key.
func (mod *Context) SetHeader(key string, val string) {
	mod.Response.Header().Set(key, val)
}

// AddHeader adds the key, value pair to the header. It appends to any existing values associated with key.
func (mod *Context) AddHeader(key string, val string) {
	mod.Response.Header().Add(key, val)
}

// Param method retrieves the parameters from url
func (mod *Context) Param(key string) string {
	return mod.Request.PathValue(key)
}

// Redirect method sets the response as a 302 redirection.
func (mod *Context) Redirect(url string) {
	mod.SetHeader("Location", url)
	mod.Status(302)
}

// Send the response immediately.
func (mod *Context) Send(code int, body any) {
	mod.Status(code)

	switch body.(type) {
	case []byte:
		mod.Response.Write(body.([]byte))
	case string:
		mod.Response.Write([]byte(body.(string)))
	case *bytes.Buffer:
		mod.Response.Write(body.(*bytes.Buffer).Bytes())
	default:
		panic(errors.New("Body type not supported"))
	}
}

// JSON sends a JSON response
func (mod *Context) JSON(code int, obj any) {
	json, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	mod.SetHeader("Content-Type", "application/json")
	mod.Send(code, json)
}

// HTML sends a HTML response
func (mod *Context) HTML(code int, html string) {
	mod.SetHeader("Content-Type", "text/html")
	mod.Send(code, html)
}
