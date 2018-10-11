package main

import "testing"

func Test_downloadFile(t *testing.T) {
	type args struct {
		url      string
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"tcp.png", args{"https://images0.cnblogs.com/blog2015/774590/201508/011703580017207.gif", "a.gif"}, false},
		{"tcp.png1", args{"ssss", "a.gif"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := downloadFile(tt.args.url, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("downloadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
