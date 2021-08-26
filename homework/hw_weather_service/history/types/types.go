package types

import "encoding/json"

// LogPostRequest ...
type LogPostRequest struct {
	Title     string `json:"title"`
	TraceID   string `json:"traceID"`
	Timestamp string `json:"timestamp"`
	Body      string `json:"body"`
}

func (l LogPostRequest) String() (s string) {
	bytes, err := json.Marshal(l)
	if err != nil {
		return
	}
	s = string(bytes)
	return
}
