package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/darjun/go-daily-lib/fyne/lib"
	"github.com/darjun/go-daily-lib/fyne/msg"
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

	// os.Unsetenv("FYNE_FONT")
}

func main() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	dirtest := "D:/omvscode/gocode_cli"
	err1 := os.MkdirAll(dirtest, 0777)
	if err1 != nil {
		fmt.Println("已经创建")
	}
	// 创建窗口
	app := app.New()

	// zt := "等待执行"

	mywin := app.NewWindow("周黑鸭pos工具集")

	l1 := widget.NewLabel("1，门店编号：")
	// l11 := widget.NewLabel(*zt)
	l2 := widget.NewLabel("1，脚本介绍：这是下载一个餐道程序")
	l3 := widget.NewButton("1，运行", func() {

		resp, err := http.Get("http://localhost:8200/pub_upload/2021-08-04/cda2jmgbav1sc5vfjs.exe")
		if err != nil {
			fmt.Println("失败")
		}
		txttest, err := os.Create("D:/omvscode/gocode_cli/餐道.exe")

		if err != nil {

			fmt.Println("报错")
		}
		io.Copy(txttest, resp.Body)
		txttest.Close()
		datapath := "D:/omvscode/gocode_cli/餐道.exe"
		cd := exec.Command(datapath)
		cd.Start()

	})

	// os.Open("text.txt")

	// 清理不能流水上传问题的脚本

	clear1 := widget.NewLabel("2,一件加速订单未上传")
	clear2 := widget.NewLabel("2,当有流水未上传的时候，请尝试一下，它会清理本地的带.ini乱码文件")

	// clear3 := widget.NewButton("运行")
	clear3 := widget.NewButton("2,运行", func() {
		// 执行订单清理业务

		lib.LibClear()
	})

	// 一键启动 餐道，vpn，pos程序

	libup1 := widget.NewLabel("3,一键启动vpn，餐道，pos程序")
	libup2 := widget.NewLabel("3,一个动作，三个程序启动")

	libup3 := widget.NewButton("3，运行", func() {
		// 请求 店务80端口是否正常
		err := lib.LibUp()

		if err {
			// msg.Dialog(mywin)
			fmt.Println("弹窗")
			// mywin1 := app.NewWindow("msg")
			msg.Dialog(mywin)

		}

	})

	// 获得扫码枪二维码图片
	libimg1 := widget.NewLabel("4,获得扫码枪设置二维码")
	libimg2 := widget.NewLabel("4,若你发现扫码抢不能自动回车，请点击运行，然后打开图片")
	libimg3 := widget.NewButton("4,运行", func() {
		test, err := lib.LibImg()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(test)

		// 打开img图片
		datapath := "D:/omvscode/gocode_cli/二维码.png"
		// img2 := canvas.NewImageFromFile(datapath)
		myimg := app.NewWindow("img")
		img2 := canvas.NewImageFromFile(datapath)
		img2.FillMode = canvas.ImageFillOriginal
		containerimg := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), img2)
		myimg.SetContent(containerimg)

		myimg.Resize(fyne.NewSize(750, 380))

		myimg.Show()
		// cd := exec.Command(datapath)
		// cd.Start()
		// fmt.Println("图片打开失败")
	})

	container1 := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), l1, l2, l3, clear1, clear2, clear3, libup1, libup2, libup3, libimg1, libimg2, libimg3)
	mywin.SetContent(container1)
	mywin.Resize(fyne.NewSize(400, 800))

	// 释放字体
	os.Unsetenv("FYNE_FONT")
	mywin.ShowAndRun()

}
