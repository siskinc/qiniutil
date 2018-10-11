package main

import (
	"context"
	"fmt"
	"os"

	"github.com/qiniu/api.v7/auth/qbox"

	"github.com/qiniu/api.v7/storage"
)

type qiniutil struct {
	accessKey string
	secretKey string
	bucket    string
	qiniuDNS  string
	config    *storage.Config
}

func (q *qiniutil) init(config *configure) (err error) {
	q.accessKey = config.AccessKey
	q.secretKey = config.SecretKey
	q.bucket = config.Bucket
	q.qiniuDNS = config.QiniuDNS
	q.config = &storage.Config{}
	q.config.UseHTTPS = false
	q.config.UseCdnDomains = false
	q.config.Zone = getStorageZone(config.Zone)
	err = q.check()
	return
}

func (q *qiniutil) check() error {
	formatErrorString := "check qiniutil is error: %s"
	if q.accessKey == "" {
		return fmt.Errorf(formatErrorString, "access key is empty string")
	}

	if q.secretKey == "" {
		return fmt.Errorf(formatErrorString, "secret key is empty string")
	}

	if q.bucket == "" {
		return fmt.Errorf(formatErrorString, "bucket is empty string")
	}

	if q.qiniuDNS == "" {
		return fmt.Errorf(formatErrorString, "qiniu dns is empty string")
	}

	if q.config == nil {
		return fmt.Errorf(formatErrorString, "config is error")
	}

	if q.config.Zone == nil {
		return fmt.Errorf(formatErrorString, "zone is error")
	}
	return nil
}

func (q *qiniutil) setAccessKey(accessKey string) {
	if accessKey != "" {
		q.accessKey = accessKey
	}
}

func (q *qiniutil) setSecretKey(secretKey string) {
	if secretKey != "" {
		q.secretKey = secretKey
	}
}

func (q *qiniutil) uploadFile(uploadFilePath, key string) (httpURL string, err error) {
	httpURL = ""
	err = nil
	// 检查文件
	if !existsPath(uploadFilePath) {
		err = fmt.Errorf("upload file path %s is not exist", uploadFilePath)
		return
	}

	if !isFile(uploadFilePath) {
		err = fmt.Errorf("upload file path %s is not a file", uploadFilePath)
		return
	}

	if key == "" {
		var uploadFile *os.File
		uploadFile, err = os.Open(uploadFilePath)
		if err != nil {
			return
		}
		key = getMd5(uploadFile)
	}

	putPolicy := storage.PutPolicy{
		Scope: q.bucket,
	}
	mac := qbox.NewMac(q.accessKey, q.secretKey)
	// 设置上传凭证有效期
	putPolicy = storage.PutPolicy{
		Scope: q.bucket,
	}
	putPolicy.Expires = 120 //示例2分钟有效期

	upToken := putPolicy.UploadToken(mac)

	formUploader := storage.NewFormUploader(q.config)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": q.bucket,
		},
	}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, uploadFilePath, &putExtra)
	if err != nil {
		return
	}
	httpURL = fmt.Sprintf("http://%s/%s\n", q.qiniuDNS, key)
	return
}
