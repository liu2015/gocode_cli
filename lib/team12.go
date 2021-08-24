package lib

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func Team12() {
	dirtest := "D:/omvscode/gocode_cli/teamviewer/TeamViewer.exe"

	cmd := exec.Command(dirtest)
	err := cmd.Start()
	if err == nil {
		return
	}
	if err != nil {
		// 下载tem12包
		resp, err := http.Get("http://172.30.10.106:8200/pub_upload/2021-08-16/cdkj626m8x3gasi8ht.zip")
		if err != nil {
			fmt.Println("下载失败")
		}

		if resp != nil {
			defer resp.Body.Close()
			// 下载包 并解压包
			resptest, err := os.Create("D:/omvscode/gocode_cli/teamviewer.zip")
			if err != nil {
				fmt.Println("是不是报错")
			}
			// 保存到本地
			io.Copy(resptest, resp.Body)
			resptest.Close()

		}

		errtest := DeCompress("D:/omvscode/gocode_cli/teamviewer.zip", "D:/omvscode/gocode_cli/")
		if errtest != nil {

			fmt.Println("解压不成功")
		}

	}

}

//解压
func DeCompress(zipFile, dest string) (err error) {
	//目标文件夹不存在则创建
	if _, err = os.Stat(dest); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dest, 0755)
		}
	}

	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}

	defer reader.Close()

	for _, file := range reader.File {
		//    log.Println(file.Name)

		if file.FileInfo().IsDir() {

			err := os.MkdirAll(dest+"/"+file.Name, 0755)
			if err != nil {
				log.Println(err)
			}
			continue
		} else {

			err = os.MkdirAll(getDir(dest+"/"+file.Name), 0755)
			if err != nil {
				return err
			}
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		filename := dest + "/" + file.Name
		//err = os.MkdirAll(getDir(filename), 0755)
		//if err != nil {
		//    return err
		//}

		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		//w.Close()
		//rc.Close()
	}
	return
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func CompressZip(src string, dest string) (err error) {

	f, err := ioutil.ReadDir(src)
	if err != nil {
		log.Println(err)
	}

	fzip, _ := os.Create(dest)
	w := zip.NewWriter(fzip)

	defer fzip.Close()
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(src + file.Name())
		if err != nil {
			log.Println(err)
		}
		_, err = fw.Write(filecontent)

		if err != nil {
			log.Println(err)
		}
		//    log.Println(n)
	}
	return
}
