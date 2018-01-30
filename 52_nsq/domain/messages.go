package domain

type Something struct {
	ID   string `json:"id"`
	TYPE string `json:"type"`
	DATA []byte `json:"data"`
}
