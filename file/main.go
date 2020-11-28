package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func main() {

}
func testBufIo() {
	var bt1b = bytes.Repeat([]byte{1}, 1)
	var bt1k = bytes.Repeat([]byte{1}, 1<<10)
	var bt3k = bytes.Repeat([]byte{1}, 1024*3)
	n := 100000
	testWrite(n, &bt1b)
	testWrite(n, &bt1k)
	testWrite(n, &bt3k)
	testBufWrite(n, &bt1b)
	testBufWrite(n, &bt1k)
	testBufWrite(n, &bt3k)
}
func testWrite(n int, data *[]byte) {
	f1, _ := ioutil.TempFile("", "testgo")
	defer f1.Close()
	t := time.Now()
	for i := 0; i < n; i++ {
		_, _ = f1.Write(*data)
	}
	log.Println(time.Since(t).Milliseconds())
}
func testBufWrite(n int, data *[]byte) {
	f1, _ := ioutil.TempFile("", "testgo")
	defer f1.Close()
	writer := bufio.NewWriterSize(f1, 1<<20)
	t := time.Now()
	for i := 0; i < n; i++ {
		_, _ = writer.Write(*data)
	}
	log.Println(time.Since(t).Milliseconds())
}
func test1() {
	isFile, _ := IsFile("./sdfsdf")
	log.Println("判断是否是文件", isFile)
	isDir, _ := IsDir("./tmpDir")
	log.Println("判断是否是目录", isDir)
	fileSize, _ := FileSize("./tmp")
	log.Println("文件的大小是", fileSize)
	// 写入文件
	FilePutContents("./t0", "t0士大夫士大夫123", 0644)
	// 读取文件
	t0Content, _ := FileGetContents("./t0")
	log.Println("to文件中的内容为", t0Content)
	// 删除文件
	Unlink("./t0")
}

// is_file()
func IsFile(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false, err
	}
	fm := fd.Mode()
	return !fm.IsDir(), nil
}

// is_dir()
func IsDir(filename string) (bool, error) {
	fd, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	fm := fd.Mode()
	return fm.IsDir(), nil
}

// filesize()
func FileSize(filename string) (int64, error) {
	info, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// file_get_contents()
func FileGetContents(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	return string(data), err
}

// unlink()
func Unlink(filename string) error {
	return os.Remove(filename)
}

// copy()
func Copy(source, dest string) (bool, error) {
	fd1, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer fd1.Close()
	fd2, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	defer fd2.Close()
	_, e := io.Copy(fd2, fd1)
	if e != nil {
		return false, e
	}
	return true, nil
}

// is_readable()
func IsReadable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_RDONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// is_writeable()
func IsWriteable(filename string) bool {
	_, err := syscall.Open(filename, syscall.O_WRONLY, 0)
	if err != nil {
		return false
	}
	return true
}

// rename()
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// touch()
func Touch(filename string) (bool, error) {
	fd, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	fd.Close()
	return true, nil
}

// mkdir()
func Mkdir(filename string, mode os.FileMode) error {
	return os.Mkdir(filename, mode)
}

// getcwd()
func Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// realpath()
func Realpath(path string) (string, error) {
	return filepath.Abs(path)
}

// basename()
func Basename(path string) string {
	return filepath.Base(path)
}

// chmod()
func Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// chown()
func Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// fclose()
func Fclose(handle *os.File) error {
	return handle.Close()
}

// filemtime()
func Filemtime(filename string) (int64, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer fd.Close()
	fileinfo, err := fd.Stat()
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}
