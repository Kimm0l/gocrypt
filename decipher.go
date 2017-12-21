package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func search() ([]string, error) {
	dir := "/Users/molina/workdir/"

	fileList := make([]string, 0)
	e := filepath.Walk(dir, func(path string, f os.FileInfo, e error) error {
		file := strings.Split(path, ".")
		if strings.EqualFold(file[len(file)-1], "aes") {
			fileList = append(fileList, path)
		}
		return e
	})

	if e != nil {
		panic(e)
	}

	fmt.Println(fileList)

	return fileList, nil
}

func decrypt(cipherstring string) string {
	// Byte array of the string
	ciphertext := []byte(cipherstring)

	// Key
	key := []byte("aserejejadejedej")

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func writeToFile(data, path string) {
	newfile := strings.Split(path, ".")
	newfile = newfile[:len(newfile)-1]
	file := strings.Join(newfile, ".")
	ioutil.WriteFile(file, []byte(data), 777)
}

func readFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func deleteFile(file string) error {
	err := os.Remove(file)
	return err
}

func main() {

	files, e := search()
	if e != nil {
		panic(e)
	}

	for _, file := range files {
		readed, e := readFromFile(file)
		if e != nil {
			panic(e)
		}
		decrypted := decrypt(string(readed))
		deleteFile(file)
		writeToFile(decrypted, file)
		fmt.Println("Descifrado")
	}
}
