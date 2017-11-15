package provisioning

import "time"

func Time() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type BaseTransport struct {
	Timestamp       int64               `json:"timestamp"`
	CustomerFid     string              `json:"customerFid"`
	SubscriptionFid string              `json:"subscriptionFid"`
	Properties      []TransportProperty `json:"properties"`
	TransportKey    string              `json:"transportKey"`
}

type Response struct {
	BaseTransport
	Type    ResponseType `json:"type"`
	Message string       `json:"message"`
	Log     []LogMessage `json:"log"`
}

type Request struct {
	BaseTransport
	Type           RequestType       `json:"type"`
	OrderFid       string            `json:"orderFid"`
	ProductFid     string            `json:"productFid"`
	PriceFid       string            `json:"priceFid"`
	StartTimestamp int64             `json:"startTimestamp"`
	RenewTimestamp int64             `json:"renewTimestamp"`
	EndTimestamp   int64             `json:"endTimestamp"`
	Configuration  map[string]string `json:"configuration"`
	Cycle          string            `json:"cycle"`
	UpdateUrl      string            `json:"updateUrl"`
}

type LogMessage struct {
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type TransportProperty struct {
	Key         string                `json:"key"`
	Type        TransportPropertyType `json:"type"`
	StringValue string                `json:"stringValue"`
	FlagValue   bool                  `json:"flagValue"`
	CountValue  int64                 `json:"countValue"`
}

type ResponseType string

const RESPONSE_SUCCESS ResponseType = "success"
const RESPONSE_PROCESSING ResponseType = "processing"
const RESPONSE_FAILED ResponseType = "failed"

type RequestType string

const REQUEST_SETUP RequestType = "setup"
const REQUEST_ACTIVATE RequestType = "activate"
const REQUEST_SUSPEND RequestType = "suspend"
const REQUEST_REACTIVATE RequestType = "reactivate"
const REQUEST_CANCEL RequestType = "cancel"
const REQUEST_END RequestType = "end"

type TransportPropertyType string

const TRANSPROP_TYPE_STRING TransportPropertyType = "string"
const TRANSPROP_TYPE_FLAG TransportPropertyType = "flag"
const TRANSPROP_TYPE_COUNT TransportPropertyType = "count"
const TRANSPROP_RETURN_TYPE_INC_COUNT TransportPropertyType = "inc.count"
const TRANSPROP_RETURN_TYPE_DEC_COUNT TransportPropertyType = "dec.count"
