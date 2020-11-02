package wreckhttp

import (
	"fmt"
	"net/http"
)

type response struct {
	Index    int
	Response http.Response
	Err      error
}

func NewResponse(index int, res http.Response, err error) *response {
	return &response{
		Index:    index,
		Response: res,
		Err:      err,
	}
}

func (r *response) GetIndex() int {
	return r.Index
}

func (r *response) SetIndex(index int) {
	r.Index = index
}

func (r *response) GetResponse() http.Response {
	return r.Response
}

func (r *response) SetResponse(res http.Response) {
	r.Response = res
}

func (r *response) GetErr() error {
	return r.Err
}

func (r *response) SetErr(err error) {
	r.Err = err
}

func (r *response) String() string {
	return fmt.Sprintf("%#v", r)
}
