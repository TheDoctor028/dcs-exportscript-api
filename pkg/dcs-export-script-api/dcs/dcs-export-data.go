package DCS

import "strconv"

type ExportData struct {
	Data map[int]string
}

func NewDcsData() ExportData {
	return ExportData{Data: map[int]string{}}
}

func (x ExportData) GetDataByUid(uid int) *string {
	item := x.Data[uid]
	if len(item) != 0 {
		return &item
	}

	return nil
}

func (x ExportData) ToString() string {
	res := ""

	for i, e := range x.Data {
		res += strconv.Itoa(i) + ":" + e + "\n"
	}
	return res
}
