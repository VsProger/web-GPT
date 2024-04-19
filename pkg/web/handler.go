package web

import (
	"Ex2_Week3/pkg/chatgpt"
	"encoding/json"
	"net/http"
)

func AskHandler(client *chatgpt.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Prompt string `json:"prompt"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error decoding request body", http.StatusBadRequest)
			return
		}

		response, err := client.AskGPT(req.Prompt)
		if err != nil {
			http.Error(w, "Failed to get response from ChatGPT", http.StatusInternalServerError)
			return
		}

		// Respond with JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"response": response})
	}
}
