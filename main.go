package main

import (
	"fmt"
	"net/url"
	"runtime"
	"time"

	"github.com/getlantern/systray"
	"github.com/atotto/clipboard"
)

var (
	previousUrl string   // Avoid parsing it over and over again
)

func main() {
	go func() {
		for x := range time.Tick(time.Second) {
			clipped, err := clipboard.ReadAll()
			if err == nil && clipped != previousUrl {
				u, invalidUrl := url.Parse(clipped)
				if invalidUrl == nil && u.Host != "" {
					// valid URL
					fmt.Printf("[%s] Processing URL: %s\n", x, clipped)
					previousUrl = processUrlItem(clipped)
					clipboard.WriteAll(previousUrl)
				}	
			}
		}
	}()
	systray.Run(onReady, onExit)
}

func onReady() {
	if runtime.GOOS == "windows" {
		// Windows won't let us have emojis. :-(
		systray.SetIcon(TrayIcon)
	} else {
		systray.SetTitle("🧹")
	}
	systray.SetTooltip("I'm cleaning URLs in the clipboard")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit cleaning URLs")

	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}

func onExit() {
	// Cleanup
	// tbd
}
