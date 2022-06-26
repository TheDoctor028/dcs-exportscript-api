package models

type DCSData struct {
	Data map[string]string
}

func NewDcsData() DCSData {
	return DCSData{Data: map[string]string{}}
}

func (x DCSData) GetDataByUid(uid string) *string {
	item := x.Data[uid]
	if len(item) != 0 {
		return &item
	}

	return nil
}
