package protocol

type (
	ExporterData struct {
		DataWithType []byte `json:"data"`
	}
	ExporterDataType byte
)

const (
	DataTypeSnapshot ExporterDataType = iota
	DataTypeRevision
	DataTypeStatus
	DataTypeStatusTransfer
)

func NewExporterData(data []byte, dataType ExporterDataType) *ExporterData {
	return &ExporterData{
		DataWithType: append(data, dataType.AsByte()),
	}
}

func (t ExporterDataType) AsByte() byte {
	return byte(t)
}
