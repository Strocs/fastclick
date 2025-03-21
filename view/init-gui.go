package view

import (
	"fastclick/settings"
	"os"
	"time"

	"fyne.io/systray"
)

const (
	YELLOW_ICON_PATH = "d:\\Documents\\Dev\\fastclick\\assets\\yellow_icon.ico"
	GREEN_ICON_PATH  = "d:\\Documents\\Dev\\fastclick\\assets\\green_icon.ico"
	MENU_TOOLTIP     = "Click to change the Toggle Key"
)

func InitSysTrayApp() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon(YELLOW_ICON_PATH))
	systray.SetTooltip(MENU_TOOLTIP)

	setToggleKeyMenu()
	systray.AddSeparator()
	setQuitMenu()
}

func addEvent(menuItem *systray.MenuItem, callback func()) {
	for range menuItem.ClickedCh {
		callback()
	}
}

func setQuitMenu() {
	mQuit := systray.AddMenuItem("Quit", "Close the application")

	go addEvent(mQuit, systray.Quit)
}

func setToggleKeyMenu() {
	currentSetting, err := settings.LoadSettings()
	if err != nil {
		panic(err)
	}

	mToggle := systray.AddMenuItem("Toggle Key: "+currentSetting.TriggerKey, MENU_TOOLTIP)

	go addEvent(mToggle, func() { onToggleKeyMenuClicked(mToggle) })
}

func onToggleKeyMenuClicked(menuItem *systray.MenuItem) {
	systray.SetIcon(getIcon(GREEN_ICON_PATH))

	menuItem.SetTooltip("Press a key")

	// Wait for the user to press a key
	time.Sleep(800 * time.Millisecond) // Simulate user interaction
	newKey := "F9"

	// Update settings with the new key
	newSetting, err := settings.UpdateSettings(newKey)
	if err != nil {
		panic(err)
	}

	menuItem.SetTitle("Toggle Key: " + newSetting.TriggerKey)
	menuItem.SetTooltip(MENU_TOOLTIP)
	systray.SetIcon(getIcon(YELLOW_ICON_PATH))
}

func onExit() {
}

func getIcon(path string) []byte {
	icon, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return icon
}
