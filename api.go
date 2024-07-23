package api

import "encoding/json"

type responseBag struct {
	HttpStatus  int         `json:"status"`
	HttpData    interface{} `json:"data"`
	HttpMessage interface{} `json:"message"`
	HttpMeta    interface{} `json:"meta"`
}

func New() *responseBag {
	return &responseBag{
		HttpStatus:  200,
		HttpData:    nil,
		HttpMessage: nil,
		HttpMeta:    nil,
	}
}

func (r *responseBag) Status(status int) *responseBag {
	r.HttpStatus = status
	return r
}

func (r *responseBag) Data(data interface{}) *responseBag {
	r.HttpData = data
	return r
}

func (r *responseBag) Message(message interface{}) *responseBag {
	r.HttpMessage = message
	return r
}

func (r *responseBag) Meta(meta interface{}) *responseBag {
	r.HttpMeta = meta
	return r
}

func (r *responseBag) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}
