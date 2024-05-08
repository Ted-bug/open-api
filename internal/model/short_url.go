package model

type ShortUrl struct {
	Id         int32  `json:"id"`
	Lurl       string `json:"lurl"`
	Surl       string `json:"surl"`
	Hash       string `json:"hash"`
	Status     int    `json:"status" gorm:"default:(-)"`
	CreateTime string `json:"create_time"`
}

func (ShortUrl) TableName() string {
	return "op_short_url"
}
