package model

type NFT struct {
	TokenID       string `json:"token_id"`
	ProductPrefix string `json:"product_prefix"`
	ProductName   string `json:"product_name"`
	ProductID     string `json:"product_id"`
	CreateTime    string `json:"create_time"`
}
