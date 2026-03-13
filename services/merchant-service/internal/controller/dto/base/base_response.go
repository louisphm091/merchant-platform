package base

type Result struct {
	ResponseCode string `json:"response_code"`
	Description  string `json:"description"`
}

type BaseResponse[T any] struct {
	RequestId       string `json:"request_id"`
	RequestDateTime string `json:"request_date_time"`
	Channel         string `json:"channel"`
	Result          Result `json:"result"`
	Data            T      `json:"data"`
}
