package api

type ResponseBag struct {
	HttpStatus  int         `json:"status"`
	HttpData    interface{} `json:"data"`
	HttpMessage interface{} `json:"message"`
	HttpMeta    interface{} `json:"meta"`
}

func New() *ResponseBag {
	return &ResponseBag{
		HttpStatus:  200,
		HttpData:    nil,
		HttpMessage: nil,
		HttpMeta:    nil,
	}
}

func (r *ResponseBag) Status(status int) *ResponseBag {
	r.HttpStatus = status
	return r
}

func (r *ResponseBag) Data(data interface{}) *ResponseBag {
	r.HttpData = data
	return r
}

func (r *ResponseBag) Message(message interface{}) *ResponseBag {
	r.HttpMessage = message
	return r
}

func (r *ResponseBag) Meta(meta interface{}) *ResponseBag {
	r.HttpMeta = meta
	return r
}
