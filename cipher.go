package main

import (
	"os"
	"net/http"
	"crypto"
	"fmt"
)


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

