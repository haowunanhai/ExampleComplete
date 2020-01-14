package ioexample

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//ioutil.ReadFile()可以非常方便的读文件，使用方不必分配空间，计算文件大小。
//将读文件抽象到最简单的 name->content的转换
func GetFileContent(name string) string {
	c, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("read file fail. Error", err)
		return ""
	}

	return string(c)
}

//ioutil.ReadAll()
func ReadFromIOReader(c string) string {
	r := strings.NewReader(c)
	s, _ := ioutil.ReadAll(r)
	fmt.Println("s:", s)
	return string(s)
}

//ioutil.WriteFile()
func WriteToFile(filename string, data []byte, perm os.FileMode) {
	ioutil.WriteFile(filename, data, perm)
	return
}

//ReadDir()
func ShowDirectory(name string) {
	s, _ := ioutil.ReadDir(name)

	for i := range s {
		fmt.Println("dir:", s[i].Name())
	}

	return
}

//TempDir()
func CreateTempDir(dir, prefix string) {
	ioutil.TempDir(dir, prefix)
	return
}

//TempFile()
func CreateTempFile(dir, prefix string) {
	ioutil.TempFile(dir, prefix)

	return
}
