package main

import (
	"github.com/atotto/clipboard"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"github.com/denisbrodbeck/machineid"
)

func main() {
	var outTE *walk.TextEdit

	MainWindow{
		Title:   "机器唯一ID",
		Size:    Size{300, 200},
		MaxSize: Size{300, 200},
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "点击获取机器唯一ID",
				OnClicked: func() {
					id, _ := machineid.ProtectedID("mid")
					clipboard.WriteAll(id)
					outTE.SetText("机器ID已复制到剪切板,可以直接粘贴使用\n\n" + id)
				},
			},
			TextEdit{AssignTo: &outTE, ReadOnly: true},
		},
	}.Run()
}
