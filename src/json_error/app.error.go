package json_error

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	e := json.NewEncoder(w).Encode(err)
	if e != nil {
		fmt.Println(e)
	}
}
