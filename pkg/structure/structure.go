package structure

type (
	Structure interface {
		Tag() Tag
		Context() string
		Timestamp() Timestamp
	}
)
