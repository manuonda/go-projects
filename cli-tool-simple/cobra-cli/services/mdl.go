package services

import (
	"github.com/culturadevops/jgt/jio"
	"os"
)

func createFolderIfNotExist(folderName string) {
	if !jio.IsFileExist(folderName) {
		jio.createDir(folderName)
	}
}

func Mdl(fileName string) {
	createFolderIfNotExist(fileName)
	MapForReplace := make(map[string]string)

}
