package model

type UniqueNum struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func (UniqueNum) TableName() string {
	return "op_unique_num"
}
