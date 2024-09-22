package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


// Post issues
func Post(owner, rep string, fields map[string]string) error {
	token := os.Getenv("GITHUB_TOKEN") // environment variable
	if token == "" {
		return fmt.Errorf("empty token")
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/", owner, rep)
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("bad request %d", resp.StatusCode)
	}
	resp.Body.Close()
	return nil
}