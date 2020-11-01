package main

import "fmt"

type Batch struct {
	Requests []Request
}

func NewBatch(requests []Request) *Batch {
	return &Batch{Requests: requests}
}

func (b *Batch) GetRequests() []Request {
	return b.Requests
}

func (b *Batch) SetRequests(requests []Request) {
	b.Requests = requests
}

func (b *Batch) String() string {
	return fmt.Sprintf("%#v", b)
}
