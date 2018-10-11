package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"

	"github.com/atotto/clipboard"
)

func main() {
	user, err := user.Current()
	homeDir := ""
	if nil == err {
		homeDir = user.HomeDir
	}
	configFilePath := flag.String("config", fmt.Sprintf("%s/.qiniutilrc.json", homeDir), "config file path")
	uploadFilePath := flag.String("f", "", "The parameter is upload file path")
	uploadFIleHTTPURL := flag.String("hf", "", "The parameter is http url of upload file")
	accessKey := flag.String("ak", "", "Qiniu Access Key")
	secretKey := flag.String("sk", "", "Qiniu Secret Key")
	key := flag.String("key", "", "file key")

	flag.Parse()

	c := &configure{}
	err = c.init(*configFilePath)
	if err != nil {
		log.Fatalf("init config error: %v\n", err)
	}

	q := &qiniutil{}
	err = q.init(c)
	if err != nil {
		log.Fatalf("init qiniutil error: %v\n", err)
	}

	q.setAccessKey(*accessKey)
	q.setSecretKey(*secretKey)

	isNeedDeleteTempFile := false
	if *uploadFilePath == "" && *uploadFIleHTTPURL != "" {
		*uploadFilePath, err = genUploadFilePathFormURL(*uploadFIleHTTPURL)
		if nil != err {
			log.Fatalf("Generate upload file path is error: %v\n", err)
		}
		err = downloadFile(*uploadFIleHTTPURL, *uploadFilePath)
		if nil != err {
			log.Fatalf("Download http file is error: %v\n", err)
		}
		isNeedDeleteTempFile = true
	}
	defer deleteHTTPUploadTempFile(isNeedDeleteTempFile, *uploadFilePath)

	httpURL, err := q.uploadFile(*uploadFilePath, *key)
	if err != nil {
		log.Fatalln(err)
	}
	if httpURL != "" {
		err = clipboard.WriteAll(httpURL)
		if nil != err {
			log.Fatalf("Copy file url to clipboard is error: %v\n", err)
		}
	}
	fmt.Printf("Upload file is success, the http URL is %s\nThe file url is in your clipboard\n", httpURL)
}
