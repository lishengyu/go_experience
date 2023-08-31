package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if args == nil || len(args) != 3 {
		fmt.Println("Usage:xxxx")
		return
	}

	srcPath := args[1]
	dstPath := args[2]
	fmt.Printf("srcPath= %s, dstPath=%s\n", srcPath, dstPath)

	if srcPath == dstPath {
		fmt.Println("src == dst")
		return
	}

	srcFile, err1 := os.Open(srcPath)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	dstFile, err2 := os.Create(dstPath)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if n == 0 {
			fmt.Println("complete")
			break
		}

		tmp := buf[:]
		dstFile.Write(tmp)
	}

	srcFile.Close()
	dstFile.Close()
}
