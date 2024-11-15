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

func (self *Http) Run(addr string) error {
	return http.ListenAndServe(addr, self.mux)
}

func (self *Http) addRoute(method string, path string, handleFunc HandleFunc) {
	self.mux.HandleFunc(method+" "+path, func(w http.ResponseWriter, r *http.Request) {
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

func (self *Http) Get(path string, handleFunc HandleFunc) {
	self.addRoute("GET", path, handleFunc)
}

func (self *Http) Post(path string, handleFunc HandleFunc) {
	self.addRoute("POST", path, handleFunc)
}

func (self *Http) Put(path string, handleFunc HandleFunc) {
	self.addRoute("PUT", path, handleFunc)
}

func (self *Http) Delete(path string, handleFunc HandleFunc) {
	self.addRoute("DELETE", path, handleFunc)
}

func (self *Http) Patch(path string, handleFunc HandleFunc) {
	self.addRoute("PATCH", path, handleFunc)
}

func (self *Http) Options(path string, handleFunc HandleFunc) {
	self.addRoute("OPTIONS", path, handleFunc)
}

func (self *Http) Head(path string, handleFunc HandleFunc) {
	self.addRoute("HEAD", path, handleFunc)
}

// Status takes an integer and sets the response statue to the integer given.
func (self *Context) Status(code int) {
	self.StatusCode = code
	self.Response.WriteHeader(code)
}

// Header gets the first value associated with the given key. If there are no values associated with
// the key. Get return "".
func (self *Context) Header(key string) string {
	return self.Request.Header.Get(key)
}

// Query method retrieves the form data.
func (self *Context) Query(key string) string {
	return self.Request.URL.Query().Get(key)
}

// SetHeader sets the header entries associated with key to the single element value. It replaces any existing values associated with key.
func (self *Context) SetHeader(key string, val string) {
	self.Response.Header().Set(key, val)
}

// AddHeader adds the key, value pair to the header. It appends to any existing values associated with key.
func (self *Context) AddHeader(key string, val string) {
	self.Response.Header().Add(key, val)
}

// Param method retrieves the parameters from url
func (self *Context) Param(key string) string {
	return self.Request.PathValue(key)
}

// Redirect method sets the response as a 302 redirection.
func (self *Context) Redirect(url string) {
	self.SetHeader("Location", url)
	self.Status(302)
}

// Send the response immediately.
func (self *Context) Send(code int, body any) error {
	self.Status(code)

	var err error
	switch body.(type) {
	case []byte:
		_, err = self.Response.Write(body.([]byte))
	case string:
		_, err = self.Response.Write([]byte(body.(string)))
	case *bytes.Buffer:
		_, err = self.Response.Write(body.(*bytes.Buffer).Bytes())
	default:
		panic(errors.New("body type not supported"))
	}

	return err
}

// JSON sends a JSON response
func (self *Context) JSON(code int, obj any) error {
	json_, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	self.SetHeader("Content-Type", "application/json")

	return self.Send(code, json_)
}

// HTML sends a HTML response
func (self *Context) HTML(code int, html string) error {
	self.SetHeader("Content-Type", "text/html")
	return self.Send(code, html)
}
