package lib

import (
	"os/exec"

	"fyne.io/fyne"
	"github.com/darjun/go-daily-lib/fyne/msg"
)

func LibStart(mywin fyne.Window) error {

	// 启动vpn网络
	vpnpath := "C:/Program Files/Sangfor/NG PDLAN/NGCtrl.exe"
	// C:\Program Files (x86)\Sangfor\NG PDLAN
	// 运行exe
	cd := exec.Command(vpnpath)
	err := cd.Start()

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
	cd3 := exec.Command(pospath)
	err3 := cd3.Start()

	if err3 != nil {
		msg.Dialog("启动pos失败", mywin)
		return err
	}

	return err3

}
