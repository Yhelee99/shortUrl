// Code generated by goctl. DO NOT EDIT.
package types

type ConvertReq struct {
	LongUrl string `json:"longUrl" validate:"required,longUrl"`
}

type ConvertResp struct {
	ShortUrl string `json:"shortUrl"`
}

type ShowReq struct {
	ShortUrl string `path:"shortUrl validata:required"`
}

type ShowResp struct {
	LongUrl string `json:"longUrl"`
}