package main

import (
	"fmt"
	"github.com/Ericwyn/GoTools/shell"
	"github.com/Ericwyn/WirelessAdbConnect/conf"
	"github.com/Ericwyn/WirelessAdbConnect/log"
	"github.com/Ericwyn/WirelessAdbConnect/ui"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	shell.Debug(true)

	// windows 下去除命令行黑窗口
	if runtime.GOOS == "windows" {
		shell.SetCmdHandler(func(cmd *exec.Cmd) {
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		})
	}

	go startHttpServer()
	ui.StartApp()
}

func startHttpServer() {
	// 手机端通过访问 http://localhost:65000/connect?ip=192.168.199.1&port=1000
	// 将无线 adb 地址和端口传送过来
	router := gin.Default()
	router.GET(
		conf.ServerApiPath,
		func(ctx *gin.Context) {
			ip := ctx.Query("ip")
			port := ctx.Query("port")

			adbConnectAddress := ip + ":" + port

			fmt.Println("get : " + adbConnectAddress)
			ctx.Status(200)

			ui.UpdateNoteLabel("尝试连接到: " + adbConnectAddress)
			res := shell.RunShellRes("adb", "connect", adbConnectAddress)
			//log.D(res)
			//if strings.Contains(res, "connect to") {
			//    ui.UpdateNoteLabel("已成功连接到 " + adbConnectAddress)
			//}
			if strings.Contains(res, "connected to") {
				go func() {
					for i := 3; i >= 0; i-- {
						ui.UpdateNoteLabel("连接成功, " + strconv.Itoa(i) + " 秒后自动关闭")
						time.Sleep(time.Second)
					}
					os.Exit(0)
				}()
			} else {
				ui.UpdateNoteLabel(res)
			}

		},
	)

	// 启动一个服务器，展示接收手机传送过来的无线数据
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(conf.ServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.D("http 服务器已启动")
	_ = s.ListenAndServe()
}
