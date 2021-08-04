package msg

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"

	"fyne.io/fyne/widget"
)

func Dialog(msg string, mymsg fyne.Window) {

	// 这是一个消息提醒弹窗

	lbmsg := widget.NewLabel("")

	// 这个是可以然输入信息的有是否项的消息弹窗
	d := dialog.NewEntryDialog("title", msg, func(s string) {
		lbmsg.SetText("这是一个消息弹窗")
	}, mymsg)
	// d := dialog.NewError("运行成功", mymsg)

	// 这个是单个消息弹窗，有一个ok按钮
	df := dialog.NewInformation("test", msg, mymsg)

	mymsg.Show()
	d.Show()
	mymsg.Show()
	df.Show()
}
