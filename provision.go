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
}

type Response struct {
	BaseTransport
	Type    string       `json:"type"`
	Message string       `json:"message"`
	Log     []LogMessage `json:"log"`
}

type Request struct {
	BaseTransport
	Type           string            `json:"type"`
	OrderFid       string            `json:"orderFid"`
	ProductFid     string            `json:"productFid"`
	PriceFid       string            `json:"priceFid"`
	StartTimestamp int64             `json:"startTimestamp"`
	RenewTimestamp int64             `json:"renewTimestamp"`
	EndTimestamp   int64             `json:"endTimestamp"`
	Configuration  map[string]string `json:"configuration"`
	Cycle          string            `json:"cycle"`
	TransportKey   string            `json:"transportKey"`
}

type LogMessage struct {
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type TransportProperty struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	StringValue string `json:"stringValue"`
	FlagValue   bool   `json:"flagValue"`
	CountValue  int64  `json:"countValue"`
}
