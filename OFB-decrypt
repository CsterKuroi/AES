package main

import (
	"fmt"
	"time"

	gfe "github.com/kiltum/go-file-encrypt"
	"github.com/prometheus/common/log"
)

func main() {
	file := "1GB.txt"
	file_pass := "123456"
	t := time.Now()
	err := gfe.DecryptFile(file+".encrypt", file_pass)
	if err != nil {
		log.Fatal("%s \n", err)
	}
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
