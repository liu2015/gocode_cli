package lib

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

// 这是一个实例，需要改造下

func LibPing() bool {

	fmt.Println("ping 172.30.90.68 店务地址")
	err := ServerPing("172.30.90.68")

	if err {
		fmt.Println("请求测试成功")
	}
	return err
}
func ServerPing(target string) bool {
	pinger, err := ping.NewPinger(target) //ping的地址，可以是www.baidu.com，也可以是127.0.0.1
	if err != nil {
		panic(err)
	}
	pinger.Count = 4              // ping的次数
	PINGTIME := time.Duration(10) // ping的时间
	pinger.Timeout = time.Duration(PINGTIME * time.Second)
	pinger.SetPrivileged(true)
	pinger.Run() // blocks until finished
	// if err1 == nil {
	// 	return true
	// }
	stats := pinger.Statistics()
	if stats.PacketsRecv >= 1 {
		fmt.Println("执行ip检测")
		fmt.Println(stats.PacketsSent)
		fmt.Println(stats.PacketLoss)
		fmt.Println(stats.IPAddr)
		fmt.Println(stats.PacketsRecv)

		return true
	}

	return false
}
