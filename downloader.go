package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)



func downloadFile(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	out, err := os.Create(fileName)
	defer out.Close()
	if err != nil {
		return err
	}
	io.Copy(out, bytes.NewReader(body))
	return nil
}
