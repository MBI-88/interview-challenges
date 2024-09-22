package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// handlePatach update issues
func handlePatch(owner, rep, number string, fields map[string]string) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("empty token")
	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, rep, number)
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", url, buf)
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

// Update update issues
func Update(owner, rep, number string, fields map[string]string) error {
	return handlePatch(owner, rep, number, fields)
}

// Close close an issues
func Close(owner, rep, number string) error {
	fields := make(map[string]string)
	fields["state"] = "open"
	return handlePatch(owner, rep, number, fields)

} 

// Reopen re-open an issues
func Reopen(owner,rep,number string) error {
	fields := map[string]string{
		"state":"open",
	}
	return handlePatch(owner,rep,number,fields)
} 
