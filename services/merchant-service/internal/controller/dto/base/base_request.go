package base

type BaseRequest[T any] struct {
	RequestId       string `json:"requestId"`
	RequestDateTime string `json:"requestDateTime"`
	Channel         string `json:"channel"`
	Data            T      `json:"data"`
}
