package structure

type Event struct {
	Tag       Tag       `json:"tag"`
	Context   string    `json:"context"`
	Timestamp Timestamp `json:"timestamp"`
}
