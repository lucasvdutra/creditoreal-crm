package health

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(Response{
		Status: "saudavel",
		Time:   time.Now().UTC(),
	})
}
