package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type configure struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket    string `json:"bucket"`
	QiniuDNS  string `json:"qiniu_dns"`
	Zone      string `json:"zone"`
}

func (c *configure) init(configFilePath string) (err error) {
	err = nil
	if !existsPath(configFilePath) {
		err = fmt.Errorf("config file %s is not exist", configFilePath)
		return
	}
	if !isFile(configFilePath) {
		err = fmt.Errorf("config file %s is not a file", configFilePath)
		return
	}

	// 载入文件
	var configFile *os.File
	configFile, err = os.Open(configFilePath)
	if err != nil {
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	for {
		if err = decoder.Decode(c); err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}
	}
	return
}
