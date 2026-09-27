package jev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JevAPI struct {
	URL    string
	Model  string
	APIKey string
	Client *http.Client
}

func NewJevAPI(url, model, apiKey string) *JevAPI {
	return &JevAPI{
		URL:    url,
		Model:  model,
		APIKey: apiKey,
		Client: &http.Client{},
	}
}

func (jev *JevAPI) SystemOne(request Request) (*Response, error) {
	request.Model = jev.Model

	body, err := json.Marshal(request)

	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		jev.URL,
		bytes.NewReader(body),
	)

	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization",	"Bearer "+jev.APIKey)

	resp, err := jev.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Message: fmt.Sprintf(
				"JEV API returned HTTP %d",
				resp.StatusCode,
			),
			Body: string(responseBody),
		}
	}

	var result Response

	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf(
			"decode JEV response: %w; body=%s",
			err,
			string(responseBody),
		)
	}

	return &result, nil
}
