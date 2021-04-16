package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string)( int64, error) {
	srcFile, err := os.OpenFile(srcFileName,os.O_RDWR,777)
	if err != nil {
		fmt.Println("err=",err)
	}
	defer srcFile.Close()
	srcFileRead := bufio.NewReader(srcFile)

	dstFile, e := os.OpenFile(dstFileName,os.O_CREATE | os.O_RDWR,777)
	if e != nil {
		fmt.Println("err=",e)
	}
	defer dstFile.Close()
	write := bufio.NewWriter(dstFile)

	return io.Copy(write,srcFileRead)

}
func main() {
	srcName := "./C824BE827326A01CA3CB6723F7E95334.pdf"
	dstName := "./newa.pdf"
	CopyFile(dstName,srcName)
	fmt.Println("hello world")
}
