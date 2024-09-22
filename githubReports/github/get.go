package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Get issues
func Get(owner, rep, number string) (*Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, rep, number)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("bad request %d", res.StatusCode)
	}
	defer res.Body.Close()
	var result Issue
	er := json.NewDecoder(res.Body).Decode(&result)
	if er != nil {
		return nil, er
	}

	return &result, nil
}
