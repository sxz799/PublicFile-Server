package model

type SystemLog struct {
	Id   int    `json:"id" gorm:"autoIncrement"`
	Msg  string `json:"msg"`
	Date string `json:"date"`
}
