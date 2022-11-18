package structure

type Event struct {
	Structure

	Tag       Tag       `json:"tag"`
	Context   string    `json:"context"`
	Timestamp Timestamp `json:"timestamp"`
}
