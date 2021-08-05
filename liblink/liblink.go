package liblink

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// 更新link指向
var Updatetest string

// 服务器上的版本
var LinkEdittest int

// 当前版本
var Lo_LinkEdit = 0

var LinkDewurtest string

var Must bool

type AutoGenerated struct {
	Code int    `json:"code"`
	Data Data   `json:"data"`
	Msg  string `json:"msg"`
}
type List struct {
	LinkID      int         `json:"link_id"`
	LinkName    string      `json:"link_name"`
	LinkURL     string      `json:"link_url"`
	LinkZip     int         `json:"link_zip"`
	LinkBrief   string      `json:"link_brief"`
	LinkContent string      `json:"link_content"`
	LinkEdit    int         `json:"link_edit"`
	LinkDewurl  string      `json:"link_dewurl"`
	LinkRemak   string      `json:"link_remak"`
	LinkAddtime interface{} `json:"link_addtime"`
	LinkOpen    int         `json:"link_open"`
}
type Data struct {
	CurrentPage int    `json:"currentPage"`
	List        []List `json:"list"`
	Total       int    `json:"total"`
}

func Liblink() bool {
	resp, err := http.Get("http://localhost:8200/module/link/list")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	var autogen AutoGenerated

	// 获得link表
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	body1 := []byte(body)

	// error1 := gconv.Struct(body1, &autogen)

	// if error1 != nil {
	// 	fmt.Println(error1)
	// }
	// 解构成结构体
	err12 := json.Unmarshal(body1, &autogen)
	if err12 != nil {
		fmt.Println(err12)
	}
	listtest := autogen.Data.List

	for _, v := range listtest {
		fmt.Println(v)

		if v.LinkName == "updata_cli" {
			Updatetest = v.LinkName
			LinkEdittest = v.LinkEdit
			LinkDewurtest = v.LinkDewurl
		}
	}

	// 对比服务器版本，若是大于则更新
	if LinkEdittest > Lo_LinkEdit {

		respd, err1 := http.Get(LinkDewurtest)

		if err1 != nil {
			fmt.Println("没有下载成功")
		}

		defer respd.Body.Close()

		dfw, err := os.Create("D:/omvscode/gocode_cli/update.exe")
		if err != nil {
			fmt.Println()
		}
		defer dfw.Close()
		// 复制到创建update文件里
		io.Copy(dfw, resp.Body)
		cd := exec.Command("D:/omvscode/gocode_cli/update.exe")
		fmt.Println("下载并开始运行")
		cd.Start()
		Must = true

	} else {
		Must = false
	}

	return Must

}
