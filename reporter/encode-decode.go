package reporter

import (
	"encoding/json"
	"net/http"
)

func CheckResponseDecoding(r *http.Response, model interface{}) {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		panic(err)
	}
}
