package structure

type Result struct {
	Tag       Tag       `json:"tag"`
	Context   string    `json:"context"`
	Timestamp Timestamp `json:"timestamp"`
}
