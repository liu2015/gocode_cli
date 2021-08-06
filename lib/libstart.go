package lib

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne"
	"github.com/darjun/go-daily-lib/fyne/msg"
)

func LibStart(mywin fyne.Window) error {

	// 启动vpn网络
	vpnpath := "C:/Program Files/Sangfor/NG PDLAN/NGCtrl.exe"
	// C:\Program Files\Sangfor\NG PDLAN\NGCtrl.exe
	// C:\Program Files (x86)\Sangfor\NG PDLAN
	// 运行exe  cmd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", Filename)
	// cd := exec.Command("cmd.exe", "/C", "start ",vpnpath)
	// Filename = "\"" + Filename + "\""
	// 也可以用转义字符在路径前后添加 \""
	// 终于搞定了，应为路径带空格，需要启动特殊的环境
	cd := exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", vpnpath)
	err := cd.Run()
	fmt.Println(err)

	if err != nil {
		msg.Dialog("启动vpn失败，请联系管理员", mywin)

	}
	// 启动餐道
	candaopath := "C:/BOMS/BOMS.exe"
	// C:\BOMS\BOMS.exe
	// // 运行exe
	cd2 := exec.Command(candaopath)
	err2 := cd2.Start()
	if err2 != nil {
		msg.Dialog("启动启动餐道失败，请联系管理员", mywin)

	}

	// 启动pos
	pospath := "D:/hisense/HiPOS6/MainExePro.exe"
	// D:\hisense\HiPOS6\MainExePro.exe
	// exec.Cmd.Dir("D:/hisense/HiPOS6/")
	// execmd := exec.Cmd{
	// 	Dir: "D:/hisense/HiPOS6/",
	// }

	cd3 := exec.Command("cmd.exe", "/C", "start "+pospath)
	// cd3.Dir("D:/hisense/HiPOS6/")
	cd3.Dir = "D:/hisense/HiPOS6/"

	// cd3.Dir("D:/hisense/HiPOS6/")
	// err3 := cd3.Start()
	err3 := cd3.Run()

	if err3 != nil {
		msg.Dialog("启动pos失败", mywin)
		return err
	}

	return err3

}
