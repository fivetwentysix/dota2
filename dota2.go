package dota2

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Dota2Request is the main struct for making
// Steam Dota2 api calls
type Dota2Request struct {
	APIBase     string
	APIEnd      string
	Method      string
	Param       string
	SteamAPIKey string
}

// NewDota2Request returns a Dota2Request struct
// to be used for API calls
func NewDota2Request(sAPI string) Dota2Request {
	return Dota2Request{
		APIBase:     "http://api.steampowered.com",
		Method:      "GET",
		SteamAPIKey: sAPI,
	}
}

// RequestSend sends the request to the specified endpoint
func (d *Dota2Request) RequestSend() (*Dota2Response, error) {
	if len(d.APIEnd) == 0 {
		return nil, errors.New("need to specify an endpoint for the API call")
	}
	url := d.APIBase + d.APIEnd + "?key=" + d.SteamAPIKey + "&language=english"
	if d.Param != "" {
		url += "&" + d.Param
	}
	req, err := http.NewRequest(d.Method, url, nil)
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var drResp Dota2Response
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&drResp)
	if err != nil {
		return nil, err
	}
	return &drResp, nil
}
