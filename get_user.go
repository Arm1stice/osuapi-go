package osuapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetUser - Get user matching specific criteria
func (api API) GetUser(options M) (*User, error) {
	// Create the request object we are going to add the query too
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/get_user", nil)
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

	// Create a list of users and unmarshal the response into it
	users := []User{}
	err = json.Unmarshal(buf, &users)
	if err != nil {
		return nil, err
	}
	// If no user was found, return nothing
	if len(users) == 0 {
		return nil, nil
	}

	user := users[0] // Get the user (because for some reason the API returns things in an array)
	// Check if the API returned an error
	if user.Error != "" {
		return nil, errors.New(user.Error)
	}
	return &user, nil
}
