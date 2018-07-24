package osuapi

import (
	"net"
	"net/http"
	"time"
)

// API - An object with an embedded API Key
type API struct {
	APIKey string
}

// M - Shorthand for map[string]string, used for options to pass to the API
type M map[string]string

// From https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout: 5 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}
var netClient = &http.Client{
	Timeout:   time.Second * 10,
	Transport: netTransport,
}

// NewAPI - Create a new API with the provided API key, returns an instance of API
func NewAPI(APIKey string) API {
	return API{
		APIKey: APIKey,
	}
}
