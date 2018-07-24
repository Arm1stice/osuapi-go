package osuapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetBeatmap - Get beatmaps matching a certain criteria
func (api API) GetBeatmap(options M) (*[]Beatmap, error) {
	// Create the request object we are going to add the query too
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/get_beatmaps", nil)
	if err != nil {
		return nil, err
	}

	// Grab the existing query
	q := req.URL.Query()
	// Add the API key to it
	q.Add("k", api.APIKey)
	// Loop through the options provided to us and add them to the query
	for k, v := range options {
		q.Add(k, v)
	}

	// Encode the querystring and add it to the request URL
	req.URL.RawQuery = q.Encode()

	// Submit a GET request
	res, err := netClient.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	// Read the response
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Create a list of beatmaps and unmarshal the response into it
	beatmaps := []Beatmap{}
	err = json.Unmarshal(buf, &beatmaps)
	if err != nil {
		return nil, err
	}

	return &beatmaps, nil
}
