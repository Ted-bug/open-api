package model

type AkSk struct {
	Id         int32  `json:"id"`
	Uid        string `json:"uid"`
	Ak         string `json:"ak"`
	Sk         string `json:"sk"`
	Show       string `json:"show"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

func (AkSk) TableName() string {
	return "op_ak_sk"
}
