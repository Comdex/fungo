package command

import (
	"fmt"
	"os"
)

//错误检测函数
func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Successful Executing!")
	}
}
