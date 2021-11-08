package libFile

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

//获得文本文件内容，以行为元素的切片
func GetFileContextLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

//获取文件内容
func GetFileContext(name string) string {
	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		//fmt.Println("Use ioutil.ReadFile to read a file:", result)
		return result
	}
	return ""
}

//拷贝文件
func CopyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//向文件插入信息
func FileInsertInfo(filename string, info string, offset int64) {
	// 打开要操作的文件 os.O_RDWR: 可读可写
	file, err := os.OpenFile(filename, os.O_RDWR, 0544)
	if err != nil {
		fmt.Printf("File open failed! err: %v\n", err)
		return
	}
	var tmpname string = "./a.tmp"
	// 新建临时文件
	tempFile, err := os.OpenFile(tmpname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Temp create failed! err: %v\n", err)
		return
	}
	// 在第二行插入新内容
	// 将原文件写入临时文件
	oldContent := GetFileContext(filename)
	upContent := oldContent[0:offset]
	downContent := oldContent[offset:]
	tempFile.WriteString(upContent)
	// 写入要插入的内容
	tempFile.WriteString(info)
	tempFile.WriteString(downContent)
	file.Close()
	tempFile.Close()
	err = os.Rename(tmpname, filename)
	if err != nil {
		fmt.Printf("Rename file raed failed! err: %v\n", err)
		return
	}
}

func CheckOrCreateDir(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, os.ModePerm)
			return err
		}
	}
	return err
}

func GetFilesFromDir(dir string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("GetFilesFromDir failed! err: %v\n", err)
		return nil, err
	}
	return files, nil
}
