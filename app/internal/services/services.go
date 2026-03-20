package services

import (
	"log"
	"main/app/internal/repository"
	"os"
)

func ShowFiles(dirPath string) []repository.File {
	var list = []repository.File{}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for i, e := range entries {
		fileInfo, err := e.Info()
		if err != nil {
			log.Fatal(err)
		}
		file := repository.File{Id: i + 1, Name: e.Name(), Size: fileInfo.Size()}
		list = append(list, file)
	}
	return list
}

func DeleteElement(slice []repository.File, index int) []repository.File {
    return append(slice[:index], slice[index+1:]...)
}