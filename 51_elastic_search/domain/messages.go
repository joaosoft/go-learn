package domain

type Something struct {
	TYPE string `json:"type"`
	ID   string `json:"id"`
	DATA []byte `json:"data"`
}
