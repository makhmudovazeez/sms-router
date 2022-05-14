package models

type SMS struct {
	id         int64  `json:"id"`
	uuid       string `json:"uuid"`
	text       string `json:"text"`
	created_at int64  `json:"created_at"`
	created_by string `json:"created_by"`
}
