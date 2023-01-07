package protocol

type (
	ExporterData struct {
		DataWithType []byte `json:"data"`
	}
	ExporterDataType byte
)

const (
	DataTypeSnapshot ExporterDataType = iota
	DataTypeStatus
	DataTypeStatusTransfer
)

func NewExporterData(data []byte, dataType ExporterDataType) *ExporterData {
	return &ExporterData{
		DataWithType: append(data, dataType.Byte()),
	}
}

func (d ExporterData) Data() []byte {
	return d.DataWithType
}

func (t ExporterDataType) Byte() byte {
	return byte(t)
}
