package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {
	putPolicy := storage.PutPolicy{
		Scope: "zeroginblog",
	}
	AccessKey := "V64rbkliOtaRUVXxNPJsvtBOftdlsv6fipN9F5SP"
	SecretKey := "IlkI7-R65M5K-2BVXL4_bAoAA_GYhDfZOsjmOcBH"
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	str := "here ok"
	read := bytes.NewBuffer([]byte(str))

	fileSize := read.Len()

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, read, int64(fileSize), &putExtra)
	if err != nil {
		fmt.Println("err:", err)
	}
	url := "http://127.0.0.1:8090" + ret.Key
	fmt.Println(url)
}
