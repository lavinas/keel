package krest

// RestCreated represents an create call return
type KRestResult struct {
	Status   int         `json:"status"`
	Title    string      `json:"title"`
	Embedded interface{} `json:"_embedded"`
}

// NewRestResult creates a new RestResult
func NewKRestResult(status int, title string, embedded interface{}) *KRestResult {
	return &KRestResult{
		Status:   status,
		Title:    title,
		Embedded: embedded,
	}
}
