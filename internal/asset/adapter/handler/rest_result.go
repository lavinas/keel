package handler

// RestCreated represents an create call return
type RestResult struct {
	Status   int         `json:"status"`
	Title    string      `json:"title"`
	Embedded interface{} `json:"_embedded"`
}

// NewRestResult creates a new RestResult
func NewRestResult(status int, title string, embedded interface{}) *RestResult {
	return &RestResult{
		Status:   status,
		Title:    title,
		Embedded: embedded,
	}
}
