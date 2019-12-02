package main

import (
	"fmt"
	"time"

	gfe "github.com/kiltum/go-file-encrypt"
	"github.com/prometheus/common/log"
)

func main() {
	//fsutil file createnew 1GB.txt 1073741824
	file := "1GB.txt"
	file_pass := "123456"
	t := time.Now()
	err := gfe.EncryptFile(file, file_pass)
	if err != nil {
		log.Fatal("%s \n", err)
	}
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
