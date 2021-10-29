package main

import (
	"runtime"
	"time"

	"github.com/markcial/track-gwei/icon"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func main() {
	runtime.LockOSThread()

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()

		data := core.NSData_WithBytes(icon.Data, uint64(len(icon.Data)))
		img := cocoa.NSImage_alloc().InitWithData__asNSImage(data)
		img.SetSize(core.Size(16, 16))
		obj.Button().SetImage(img)
		obj.Button().SetImagePosition(cocoa.NSImageLeft)
		obj.Button().SetTitle("Price: " + getGwei())

		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		ticker := time.NewTicker(10 * time.Second)

		go func() {
			for {
				select {
				case <-ticker.C:
					obj.Button().SetTitle("Price: " + getGwei())
				}
			}
		}()

		menu := cocoa.NSMenu_New()
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)
	})
	app.Run()
}
