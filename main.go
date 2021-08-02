package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	fmt.Println("=============")

	// os.Unsetenv("FYNE_FONT")
}

func main() {

	fmt.Println("你好，这是一个运行程序")

	app := app.New()

	mywin := app.NewWindow("周黑鸭pos工具集")
	mywin.SetContent(widget.NewLabel("门店编号：   \n"))

	l1 := widget.NewLabel("门店编号：")
	l2 := widget.NewLabel("脚本介绍：这是下载一个餐道程序")
	l3 := widget.NewButton("运行", func() {

		resp, err := http.Get("https://www.baidu.com/")
		if err != nil {
			fmt.Println("失败")
		}
		txttest, err := os.Create("text.txt")
		io.Copy(txttest, resp.Body)

	})

	// os.Open("text.txt")

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), l1, l2, l3)
	mywin.SetContent(container)
	mywin.Resize(fyne.NewSize(400, 700))

	// 释放字体
	os.Unsetenv("FYNE_FONT")
	mywin.ShowAndRun()

}
