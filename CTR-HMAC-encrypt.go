package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const BUFFER_SIZE int = 4096
const IV_SIZE int = 16

func encrypt(in io.Reader, out io.Writer, keyAes, keyHmac []byte) (err error) {

	iv := make([]byte, IV_SIZE)
	_, err = rand.Read(iv)
	if err != nil {
		return err
	}

	aes, err := aes.NewCipher(keyAes)
	if err != nil {
		return err
	}

	ctr := cipher.NewCTR(aes, iv)
	hmac := hmac.New(sha256.New, keyHmac)

	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := in.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		outBuf := make([]byte, n)
		ctr.XORKeyStream(outBuf, buf[:n])
		hmac.Write(outBuf)
		out.Write(outBuf)

		if err == io.EOF {
			break
		}
	}

	out.Write(iv)
	hmac.Write(iv)
	out.Write(hmac.Sum(nil))

	return nil
}

//
// For the demo
//

type devZero byte

func (z devZero) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = byte(z)
	}
	return len(b), nil
}

func mockDataSrc(size int64) io.Reader {
	var z devZero
	return io.LimitReader(z, size)
}

func main() {
	//fsutil file createnew 1GB.txt 1073741824
	filePathIn := "1GB.txt"
	filePathOut := "1GB.txt.enc"
	inFile, err := os.Open(filePathIn)
	if err != nil { log.Fatal("%s \n", err) }
	defer inFile.Close()

	outFile, err := os.Create(filePathOut)
	if err != nil { log.Fatal("%s \n", err) }
	defer outFile.Close()

	keyAes, _ := hex.DecodeString("6368616e676520746869732070617373")
	keyHmac := keyAes // don't do this

	t := time.Now()
	err = encrypt(inFile, outFile, keyAes, keyHmac)

	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(t)
	fmt.Println("app elapsed:", elapsed)
}
