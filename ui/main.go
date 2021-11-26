package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Ericwyn/WirelessAdbConnect/conf"
	"github.com/Ericwyn/WirelessAdbConnect/log"
	"github.com/Ericwyn/WirelessAdbConnect/ui/resource"
	"github.com/skip2/go-qrcode"
	"net"
	"strconv"
)

var mainWindow fyne.Window
var mainApp fyne.App
var connectLog *widget.Label

func StartApp() {
	// 设置整个 app 的信息
	mainApp = app.New()
	mainApp.SetIcon(resource.ResourceIcon)
	mainApp.Settings().SetTheme(&resource.CustomerTheme{})

	showAboutUi()
}

func showAboutUi() {
	mainWindow = mainApp.NewWindow("WirelessAdb")

	mainWindow.Resize(fyne.Size{
		Width: 350,
		//Height: 200,
	})
	mainWindow.CenterOnScreen()

	serverAddress := "http://" + getIpAddress() + ":" + strconv.Itoa(conf.ServerPort) + conf.ServerApiPath

	qrCode := canvas.NewImageFromResource(generalQrCodeResource(serverAddress))
	qrCode.SetMinSize(fyne.Size{
		Width:  256,
		Height: 256,
	})

	connectLog = widget.NewLabel("未检测到连接")
	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel(""),
		container.NewCenter(widget.NewLabel("请扫描以下二维码进行连接无线 adb")),
		//widget.NewLabel(""),
		container.NewCenter(qrCode),
		container.NewCenter(widget.NewLabel(serverAddress)),
		widget.NewLabel(""),
		container.NewCenter(connectLog),
		widget.NewLabel(""),
	))

	mainWindow.SetOnClosed(func() {

	})

	mainWindow.ShowAndRun()
}

func generalQrCodeResource(msg string) *fyne.StaticResource {
	var png []byte
	png, _ = qrcode.Encode(msg, qrcode.Medium, 256)
	return &fyne.StaticResource{
		StaticName:    "qrcode",
		StaticContent: png,
	}
}

func getIpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.E(err)
	}
	var ip string = "localhost"
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip
			}
		}
	}
	return ip
}

func UpdateNoteLabel(msg string) {
	if connectLog != nil {
		connectLog.SetText(msg)
		connectLog.Refresh()
	}
}
