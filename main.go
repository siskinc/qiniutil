package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
)

func main() {
	user, err := user.Current()
	homeDir := ""
	if nil == err {
		homeDir = user.HomeDir
	}
	configFilePath := flag.String("config", fmt.Sprintf("%s/.qiniutilrc.json", homeDir), "config file path")
	uploadFilePath := flag.String("file", "", "The paramater is upload file path")
	accessKey := flag.String("access_key", "", "Qiniu Access Key")
	secretKey := flag.String("secret_key", "", "Qiniu Secret Key")
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

	httpURL, err := q.uploadFile(*uploadFilePath, *key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Upload file is success, the http URL is %s\n", httpURL)
}
