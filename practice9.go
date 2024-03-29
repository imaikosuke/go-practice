package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// アプリケーションの作成
	myApp := app.New()
	// ウィンドウの作成
	myWindow := myApp.NewWindow("Hello Fyne")

	// ボタンの作成。クリックされたときの動作を定義
	helloButton := widget.NewButton("Say Hello", func() {
		dialog.ShowInformation("Hello", "Hello, Fyne!", myWindow)
	})

	// ボタンをウィンドウに追加
	myWindow.SetContent(container.NewVBox(helloButton))

	// ウィンドウの表示とアプリケーションの起動
	myWindow.ShowAndRun()
}
