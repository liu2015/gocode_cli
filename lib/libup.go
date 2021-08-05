package lib

import (
	"fmt"
)

func LibUp() bool {

	fmt.Println("检查网络")

	err := LibPing()

	if err {
		fmt.Println(err)
	}
	return err

}
