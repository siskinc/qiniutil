package main

import (
	"testing"
)

func Test_checkURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid url for http", args{"http://www.baidu.com"}, true},
		{"valid url for https", args{"https://www.baidu.com"}, true},
		{"not valid url", args{"htps://www.baidu.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkURL(tt.args.url); got != tt.want {
				t.Errorf("checkURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genUploadFilePathFormURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"true", args{"http://www.baidu.com/a.jpg"}, "a.jpg", false},
		{"false", args{"htp://www.baidu.com/a.jp"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genUploadFilePathFormURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("genUploadFilePathFormURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("genUploadFilePathFormURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteHTTPUploadTempFile(t *testing.T) {
	type args struct {
		need bool
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"first", args{true, "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteHTTPUploadTempFile(tt.args.need, tt.args.path)
		})
	}
}
