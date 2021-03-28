package responses

import (
	"encoding/json"
	"net/http"
)

func Return_json_response(w http.ResponseWriter, response_shape interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response_shape)
}
