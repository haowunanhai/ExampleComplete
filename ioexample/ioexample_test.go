package ioexample

import (
	"os"
	"testing"
)

func TestIo(t *testing.T) {
	//读文件
	name := "test1"
	t.Log("filename", name, "content:", GetFileContent(name))

	//读io.Reader
	content := "Hello world"
	t.Log("content:", ReadFromIOReader(content))

	//写文件
	WriteToFile(name, []byte("Helloworld written"), os.ModeAppend)
	t.Log("filename", name, "after written content:", GetFileContent(name))

	//读目录
	ShowDirectory("/data0/liubin11/home/")

	//创建随机文件
	CreateTempFile("/data0/liubin11/ExampleComplete/ioexample", "LB")

	//创建随机目录
	CreateTempDir("/data0/liubin11/ExampleComplete/ioexample", "LB")
}
