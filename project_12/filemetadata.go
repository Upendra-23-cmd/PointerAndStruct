package main

import (
	"fmt"
	"os"
)


type Information struct{
	Size  int32
	Type  string
	Perm  string
}

func getMetaData(path string)(*Information, error){

	info, err := os.Stat(path)
	if err != nil {
		return nil,err
	}

	// Determine file type
	FileType := "file"
	if info.IsDir(){
		FileType = "Directory"
	}

	// Convert permission to string (like rwxr-xr-x)
	perm := info.Mode().Perm().String()

	return &Information{
		Size: int32(info.Size()),
		Type: FileType,
		Perm: perm,
	}, nil

}
func main() {
	meta, err := getMetaData("Readme.md")
	if err != nil {
		fmt.Println("error while getting the filedata : ", err)
		return
	}
	fmt.Printf("Metadata of the file is \n %+v\n",meta)
}
