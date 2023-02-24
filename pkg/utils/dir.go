package utils

import (
	"io/fs"
	"os"
)

func ExistDir(dir string) (exist bool) {
	_, err := os.Stat(dir)
	if err != nil {
		return
	}
	exist = true
	return
}

func CreateDir(dir string) error {
	exist := ExistDir(dir)
	if exist {
		return nil
	}
	err := os.MkdirAll(dir, fs.ModePerm)
	return err
}
