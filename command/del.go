package command

import (
	//"fmt"
	"os"
)

func DelCommand(dirName string) {
	err := os.RemoveAll(dirName)
	if err != nil {

	}
	CheckError(err)
}
