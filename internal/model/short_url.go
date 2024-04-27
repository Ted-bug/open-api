package model

type ShortUrl struct {
	Id         int32  `json:"id"`
	Url        string `json:"url"`
	Short      string `json:"short"`
	Hash       string `json:"hash"`
	Status     int    `json:"status"`
	CreateTime string `json:"create_time"`
}

func (ShortUrl) TableName() string {
	return "op_short_url"
}
