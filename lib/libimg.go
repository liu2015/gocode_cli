package lib

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func LibImg() (testing string, err error) {

	fmt.Println("运行并打开图片")

	url := "http://localhost:8200/pub_upload/2021-08-04/cdaellx3015w9zua7p.png"

	resp, err := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("下载失败")
	}
	// resp.Body.Close()

	fie, err := os.Create("G:/omvscode/gocode_cli/二维码.png")
	if err != nil {
		fmt.Println("创建失败")
	}

	io.Copy(fie, bytes.NewReader(body))
	defer fie.Close()

	defer resp.Body.Close()

	// 打开img图片
	// datapath := "G:/omvscode/gocode_cli/二维码.png"
	// cd := exec.Command(datapath)
	// cd.Start()
	// fmt.Println("图片打开失败")

	return "下载图片完成", nil
}
