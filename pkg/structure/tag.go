package structure

type (
	Tag string
)

func (t Tag) String() string {
	return string(t)
}

func FromString(tagStr string) Tag {
	return Tag(tagStr)
}

// TODO 这里边插桩边加具体的
