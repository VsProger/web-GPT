package chatgpt

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	APIKey     string
	APIBaseURL string
}

type APIRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type APIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		APIKey:     apiKey,
		APIBaseURL: baseURL,
	}
}

func (c *Client) AskGPT(prompt string) (string, error) {
	requestData := map[string]interface{}{
		"prompt":      prompt,
		"max_tokens":  150,
		"temperature": 0.7,
	}
	requestDataBytes, err := json.Marshal(requestData)
	if err != nil {
		log.Printf("Failed to marshal request body: %v", err)
		return "", err
	}

	request, err := http.NewRequest("POST", c.APIBaseURL, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		log.Printf("Failed to create HTTP request: %v", err)
		return "", err
	}

	request.Header.Set("Authorization", "Bearer "+c.APIKey)
	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return "", err
	}
	defer response.Body.Close()

	var responseContent map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseContent); err != nil {
		log.Printf("Failed to decode response body: %v", err)
		return "", err
	}

	log.Printf("API response: %+v", responseContent)

	if choices, ok := responseContent["choices"].([]interface{}); ok && len(choices) > 0 {
		if choiceMap, ok := choices[0].(map[string]interface{}); ok {
			if text, ok := choiceMap["text"].(string); ok {
				return text, nil
			}
		}
	}
	return "No response from GPT", nil
}
