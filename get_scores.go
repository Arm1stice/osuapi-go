package osuapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetScores - Get top scores for a beatmap
func (api API) GetScores(options M) (*[]Score, error) {
	// Create the request object we are going to add the query too
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/get_scores", nil)
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

	// Create a list of scores and unmarshal the response into it
	scores := []Score{}
	err = json.Unmarshal(buf, &scores)
	if err != nil {
		return nil, err
	}

	return &scores, nil
}
