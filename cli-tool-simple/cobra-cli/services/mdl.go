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
	MapForReplace["%MODELNAME%"] := fileName
	newName := "models/" + fileName + ".go"
	templateName :=	"/home/manuonda/projects/go-projects/cli-tool-simple/cobra-cli/template/model.stub"
	jio.NewFileforTemplate(newName, templateName, MapForReplace)

}
