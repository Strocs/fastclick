package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

func InitGUI() {
	a := app.New()
	initSystemTray(a)
	a.Run()
}

func initSystemTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		menu := createMenu("FastClick", []*fyne.MenuItem{
			addMenuItem("Toggle Key", openToggleKeyModal),
		})
		desk.SetSystemTrayMenu(menu)
	}
}

func createMenu(name string, items []*fyne.MenuItem) *fyne.Menu {
	return fyne.NewMenu(name, items...)
}

func addMenuItem(name string, action func()) *fyne.MenuItem {
	return fyne.NewMenuItem(name, action)
}

func openToggleKeyModal() {

}
