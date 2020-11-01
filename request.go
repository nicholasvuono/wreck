package main

import (
	"fmt"
	"net/url"
)

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    map[string]string
}

func NewRequest(method string, URL string, headers map[string]string, body map[string]string) *Request {
	_, err := url.Parse(URL)
	if err != nil {
		explain(err)
		return nil
	}
	return &Request{
		Method:  method,
		URL:     URL,
		Headers: headers,
		Body:    body,
	}
}

func (r *Request) GetMethod() string {
	return r.Method
}

func (r *Request) SetMethod(method string) {
	r.Method = method
}

func (r *Request) GetURL() string {
	return r.URL
}

func (r *Request) SetURL(URL string) {
	r.URL = URL
}

func (r *Request) GetHeaders() map[string]string {
	return r.Headers
}

func (r *Request) SetHeaders(headers map[string]string) {
	r.Headers = headers
}

func (r *Request) GetBody() map[string]string {
	return r.Body
}

func (r *Request) SetBody(body map[string]string) {
	r.Body = body
}

func (r *Request) String() string {
	return fmt.Sprintf("%#v", r)
}
