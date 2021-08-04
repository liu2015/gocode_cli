package msg

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"

	"fyne.io/fyne/widget"
)

func Dialog(mymsg fyne.Window) {

	// 这是一个消息提醒弹窗

	lbmsg := widget.NewLabel("这是一个消息提醒")

	d := dialog.NewEntryDialog("title", "运行成功", func(s string) {
		lbmsg.SetText("这是一个消息弹窗")
	}, mymsg)
	// d := dialog.NewError("运行成功", mymsg)
	mymsg.Show()
	d.Show()
	mymsg.Show()
}
