package util
import(
	"github.com/culturadevops/jgt/jio"
)

/**
 * Permite verificar si existe un file 
 * determinado
**/
func IsFileExist(nameFile *string) bool{
   var exist
   if _, err := os.Stat("sample.txt"); err == nil {
      return false
	} 
	return true
}

func createDirectory(directoy string) {
	if err := os.Mkdir(directory , os.ModePerm); err != nil {
        log.Fatal(err)
    }

}

func createFolderIfNotExist(fileName string){
   if(!fileExist(fileName)) {
	 createDirectory(fileName)
   }
}
