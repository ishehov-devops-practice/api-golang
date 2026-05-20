package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GraphQLRequest struct {
	Query string `json:"query"`
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method != http.MethodPost {
			http.Error(w, "GraphQL requests must be POST", http.StatusMethodNotAllowed)
			return
		}

		var req GraphQLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Mocking a successful GraphQL payload response
		responseData := map[string]interface{}{
			"data": map[string]interface{}{
				"status": "Go GraphQL API is functional",
				"echo":   req.Query,
			},
		}

		json.NewEncoder(w).Encode(responseData)
	})

	fmt.Println("Go GraphQL server listening on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}