package utils

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Qiniu struct {
	AccessKey  string
	SecretKey  string
	Bucket     string
	ReturnBody string
}

func NewQiniu() *Qiniu {
	return &Qiniu{
		AccessKey:  "ui_qeUFkfc-hqgdLpS43ImTd3maOiloStVNGH4yV",
		SecretKey:  "QTpvj74GxuHB3_A5cx0wZ4QYJa_w-W4ieIdAcXl_",
		Bucket:     "inwind-blog",
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"`,
	}
}

func (q *Qiniu) UpLoadToken() string {
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)

	putPolicy := storage.PutPolicy{
		Scope:      q.Bucket,
		Expires:    7200, //有效期
		ReturnBody: q.ReturnBody,
	}
	return putPolicy.UploadToken(mac)
}
