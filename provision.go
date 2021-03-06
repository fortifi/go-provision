package provisioning

import (
	"time"
	"crypto/md5"
	"fmt"
	"strings"
	"errors"
)

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
	VerifyHash     string            `json:"verifyHash"`
}

func (r Request) IsVerified(provisioningKey string) bool {
	return strings.ToLower(r.VerifyHash) == fmt.Sprintf("%x", md5.Sum([]byte(r.TransportKey+provisioningKey)))
}

func (r Request) Property(propertyKey string) (*TransportProperty, error) {
	for _, prop := range r.Properties {
		if prop.Key == propertyKey {
			return &prop, nil
		}
	}
	return nil, errors.New("property not set")
}

type LogMessage struct {
	Timestamp int64   `json:"timestamp"`
	Type      LogType `json:"type"`
	Message   string  `json:"message"`
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

type LogType string

const LOG_DEBUG LogType = "debug"
const LOG_INFO LogType = "info"
const LOG_WARNING LogType = "warning"
const LOG_ERROR LogType = "error"
const LOG_SUCCESS LogType = "success"
