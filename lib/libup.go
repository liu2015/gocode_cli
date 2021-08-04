package lib

import (
	"fmt"
)

func LibUp() bool {

	fmt.Println("一键启动三个程序")

	err := LibPing()

	if err {
		fmt.Println(err)
	}
	return err

}
