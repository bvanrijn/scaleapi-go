package scaleapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// Client represents a client that interacts with the Scale API
type Client struct {
	Key string
}

// CreateCategorizationTask creates a new categorization task
func (client *Client) CreateCategorizationTask(callbackURL, instruction, attachmentType, attachment string, categories []string) (task Task, err error) {
	v := url.Values{}
	v.Set("callback_url", callbackURL)
	v.Set("instruction", instruction)
	v.Set("attachment_type", attachmentType)
	v.Set("attachment", attachment)

	for _, category := range categories {
		v.Add("categories", category)
	}

	body := strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", "https://api.scaleapi.com/v1/task/categorize", body)
	if err != nil {
		return task, err
	}
	req.SetBasicAuth(client.Key, "")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&task)

	return task, err
}

// FetchTask retrieves a specific task.
func (client *Client) FetchTask(taskID string) (task Task, err error) {
	req, err := http.NewRequest("GET", "https://api.scaleapi.com/v1/task/"+taskID, nil)
	if err != nil {
		return task, err
	}
	req.SetBasicAuth(client.Key, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&task)

	return task, err
}

// CancelTask cancels a given task
func (client *Client) CancelTask(taskID string) (task Task, err error) {
	req, err := http.NewRequest("GET", "https://api.scaleapi.com/v1/task/"+taskID+"/cancel", nil)
	if err != nil {
		return task, err
	}
	req.SetBasicAuth(client.Key, "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return task, err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&task)

	return task, err
}
