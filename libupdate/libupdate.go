package libupdate

import (
	"fmt"

	"github.com/darjun/go-daily-lib/fyne/liblink"
)

func LibUpdate() bool {

	Mettf := liblink.Liblink()
	fmt.Println("检测版本更新", Mettf)

	return Mettf
}
