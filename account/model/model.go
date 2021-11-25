package model

type Account struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phonenumber int64  `json:"phonenumber"`
	Balance     int64  `json:"balance"`
}
