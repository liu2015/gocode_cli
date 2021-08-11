package lib

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne"
	"github.com/darjun/go-daily-lib/fyne/msg"
)

func LibClear(mywin fyne.Window) {

	fmt.Println("运行订单上传清理")

	// 运行cmd,然后启动脚本，一键清理门店未上传数据 D:\hisense\HiPOS6
	url := "D:/hisense/HiPOS6/ini脚本.bat"

	cd := exec.Command("cmd.exe", "/C", "start "+url)
	cd.Dir = "D:/hisense/HiPOS6/"

	err := cd.Run()

	if err != nil {
		msg.Dialog("启动脚本失败，请检查D:/hisense/HiPOS6/是否有脚本", mywin)
	} else {

		// 启动pos通讯exe
		url2 := "D:/hisense/HiPOS6/HiCTP20.exe"

		cd1 := exec.Command("cmd.exe", "/C", "start "+url2)
		cd.Dir = "D:/hisense/HiPOS6/"

		err1 := cd1.Run()

		if err1 != nil {
			msg.Dialog("启动脚本失败，请检查D:/hisense/HiPOS6/是否有脚本", mywin)
		}
	}
}
