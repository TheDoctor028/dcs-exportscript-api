package DCS

type ExportData struct {
	Data map[string]string
}

func NewDcsData() ExportData {
	return ExportData{Data: map[string]string{}}
}

func (x ExportData) GetDataByUid(uid string) *string {
	item := x.Data[uid]
	if len(item) != 0 {
		return &item
	}

	return nil
}
