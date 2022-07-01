package models

type DCSExportData struct {
	Data map[string]string
}

func NewDcsData() DCSExportData {
	return DCSExportData{Data: map[string]string{}}
}

func (x DCSExportData) GetDataByUid(uid string) *string {
	item := x.Data[uid]
	if len(item) != 0 {
		return &item
	}

	return nil
}
