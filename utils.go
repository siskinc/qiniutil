package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"

	"github.com/qiniu/api.v7/storage"
)

func existsPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func isFile(path string) bool {
	return !isDir(path)
}

func getStorageZone(zone string) *storage.Zone {
	switch zone {
	case "Huadong":
		return &storage.ZoneHuadong
	case "Huabei":
		return &storage.ZoneHuabei
	case "Huanan":
		return &storage.ZoneHuanan
	case "Beimei":
		return &storage.ZoneBeimei
	case "Xinjiapo":
		return &storage.ZoneXinjiapo
	default:
		return nil
	}
}

func getMd5(file io.Reader) string {
	md5 := md5.New()
	io.Copy(md5, file)
	return hex.EncodeToString(md5.Sum(nil))
}
