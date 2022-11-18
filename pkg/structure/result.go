package structure

type Result struct {
	Structure

	Tag       Tag       `json:"tag"`
	Context   string    `json:"context"`
	Timestamp Timestamp `json:"timestamp"`
}
