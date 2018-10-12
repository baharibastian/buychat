package respond

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func RespondWithError(w http.ResponseWriter, code int, message []string) {
	f := map[string]interface{}{
		"Message" : message,
	}
	RespondWithJSON(w, code, f)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	fmt.Println(payload)
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
